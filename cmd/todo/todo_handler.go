package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/markphelps/optional"
	"io"
	"os"
)

func LastId(dstFileName string) (int, error) {
	fi, err := os.OpenFile(dstFileName, os.O_RDONLY, 0644)
	if err != nil {
		fi, err = os.Create(dstFileName)
		if err != nil {
			return -1, err
		}
	}
	defer func() {
		err := fi.Close()
		if err != nil {
			panic(err)
		}
	}()

	var todos []Todo
	err = json.NewDecoder(fi).Decode(&todos)
	if err != nil {
		return 0, nil
	}

	id := 0
	for _, v := range todos {
		if v.ID > id {
			id = v.ID
		}
	}

	return id, nil
}

func Add(todo Todo, dstFileName string) error {
	fi, err := os.OpenFile(dstFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fi, err = os.Create(dstFileName)
		if err != nil {
			return err
		}
	}
	defer func() {
		err := fi.Close()
		if err != nil {
			panic(err)
		}
	}()
	var todos []Todo

	err = json.NewDecoder(fi).Decode(&todos)
	if err != nil && err != io.EOF {
		return err
	}

	todos = append(todos, todo)

	_, err = fi.Seek(0, 0)
	if err != nil {
		return err
	}

	err = fi.Truncate(0)
	if err != nil {
		return err
	}

	err = json.NewEncoder(fi).Encode(&todos)
	if err != nil {
		return err
	}

	return nil
}

func MakeDone(id int, dstFileName string) error {
	lastId, err := LastId(dstFileName)
	if err != nil {
		return err
	}

	if id > lastId {
		return errors.New("this id doesn't exist")
	}

	fi, err := os.OpenFile(dstFileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err := fi.Close()
		if err != nil {
			panic(err)
		}
	}()

	var todos []Todo
	err = json.NewDecoder(fi).Decode(&todos)
	if err != nil {
		return err
	}

	for i, v := range todos {
		if v.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	_, err = fi.Seek(0, 0)
	if err != nil {
		return err
	}

	err = fi.Truncate(0)
	if err != nil {
		return err
	}

	err = json.NewEncoder(fi).Encode(&todos)
	if err != nil {
		return err
	}

	return nil
}

func Update(id int, dstFileName string, desc, state optional.String) error {
	lastId, err := LastId(dstFileName)
	if err != nil {
		return err
	}

	if id > lastId {
		return errors.New("this id doesn't exist")
	}

	fi, err := os.OpenFile(dstFileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err := fi.Close()
		if err != nil {
			panic(err)
		}
	}()

	var todos []Todo
	err = json.NewDecoder(fi).Decode(&todos)
	if err != nil {
		return err
	}

	for i, v := range todos {
		if v.ID == id {
			val, err := desc.Get()
			if err == nil {
				newTodo := NewTodo(v.ID, v.CreatedAt, val)
				todos[i] = newTodo
				break
			}

			st, err := state.Get()
			if err == nil {
				switch st {
				case "in-process":
					todos[i].SetState(InProcess)
					break
				case "done":
					err := MakeDone(v.ID, dstFileName)
					if err != nil {
						return err
					}

					return nil
				default:
					panic(fmt.Errorf("unknown state: %s", st))
				}

			} else {
				panic("Missing arguments in function: todo.Todo.Update")
			}
		}
	}

	_, err = fi.Seek(0, 0)
	if err != nil {
		return err
	}

	err = fi.Truncate(0)
	if err != nil {
		return err
	}

	err = json.NewEncoder(fi).Encode(&todos)
	if err != nil {
		return err
	}

	return nil
}

func List(dstFileName string) ([]Todo, error) {
	fi, err := os.OpenFile(dstFileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := fi.Close()
		if err != nil {
			panic(err)
		}
	}()

	var todos []Todo
	err = json.NewDecoder(fi).Decode(&todos)
	if err != nil {
		return nil, err

	}

	return todos, nil
}
