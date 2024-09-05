package todo

import (
	"encoding/json"
	"os"
	"time"

	"github.com/markphelps/optional"
)

var todoId = 0

type State string

const (
	Added     State = "Added"
	InProcess State = "In process"
	Done      State = "Done"
)

func StateToString(st State) string {
	switch st {
	case Added:
		return "Added"
	case InProcess:
		return "InProcess"
	case Done:
		return "Done"
	default:
		return ""
	}
}

type Todo struct {
	createdAt   time.Time
	updatedAt   time.Time
	id          int
	state       State
	description string
}

func NewTodo(createdAt time.Time, description string) Todo {
	todoId++
	return Todo{
		createdAt:   createdAt,
		updatedAt:   time.Now(),
		id:          todoId,
		state:       Added,
		description: description,
	}
}

func Add(jsonData []byte, dstFileName string) {
	fi, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	err = json.NewEncoder(fi).Encode(jsonData)
	if err != nil {
		panic(err)
	}
}

func MakeDone(id int, dstFileName string) {
	fi, err := os.OpenFile(dstFileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	var todos []Todo
	err = json.NewDecoder(fi).Decode(&todos)
	if err != nil {
		panic(err)
	}

	for i, v := range todos {
		if v.id == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	fi, err = os.Create(dstFileName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	err = json.NewEncoder(fi).Encode(todos)
	if err != nil {
		panic(err)
	}
}

func Update(id int, dstFileName string, desc, state optional.String) {
	fi, err := os.OpenFile(dstFileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	var todos []Todo
	err = json.NewDecoder(fi).Decode(&todos)
	if err != nil {
		panic(err)
	}

	for i, v := range todos {
		if v.id == id {
			val, descErr := desc.Get()
			st, stErr := state.Get()

			if descErr == nil {
				newTodo := NewTodo(v.createdAt, val)
				todos[i] = newTodo
				break

			} else if stErr == nil {
				newTodo := NewTodo(v.createdAt, v.description)

				switch st {
				case "In progress":
					newTodo.SetState(InProcess)
					break
				case "Done":
					MakeDone(newTodo.id, dstFileName)
					return
				}

			} else {
				panic("Missing arguments in function: todo.Todo.Update")
			}
		}
	}

	fi, err = os.Create(dstFileName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	err = json.NewEncoder(fi).Encode(todos)
	if err != nil {
		panic(err)
	}
}

func (t *Todo) List(dstFileName string) []Todo {
	fi, err := os.OpenFile(dstFileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	var todos []Todo
	err = json.NewDecoder(fi).Decode(&todos)
	if err != nil {
		panic(err)
	}

	return todos
}

func (t *Todo) SetState(st State) {
	t.state = st
}
