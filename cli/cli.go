package cli

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
	"todo_cli/main/todo"

	"github.com/markphelps/optional"
)

type Command string

const (
	New       Command = "new"
	Update    Command = "udpate"
	InProcess Command = "in-process"
	Done      Command = "done"
	List      Command = "list"
)

func CommandFromString(s string) (Command, error) {
	switch s {
	case "new":
		return New, nil
	case "update":
		return Update, nil
	case "in-process":
		return InProcess, nil
	case "done":
		return Done, nil
	default:
		return "", errors.New("Invalid command")
	}
}

type Config struct {
	command Command
	id      optional.Int64
	desc    optional.String
}

func NewConfig() (Config, error) {
	args := os.Args
	if len(args) == 1 || len(args) > 4 {
		return Config{}, errors.New("Missing args")
	}

	command, err := CommandFromString(args[1])
	if err != nil {
		return Config{}, err
	}

	var desc optional.String
	var id optional.Int64
	if command != New {
		var tryId, err = strconv.ParseInt(args[2], 0, 32)
		id = optional.NewInt64(tryId)
		if err != nil {
			return Config{}, err
		}

		if command == Update {
			desc = optional.NewString(args[3])
		}
	}

	return Config{
		command: command,
		id:      id,
		desc:    desc,
	}, nil
}

func (c *Config) Handle() error {
	switch c.command {
	case New:
		todoDesc, err := c.desc.Get()
		if err != nil {
			panic(err)
		}

		newTodo := todo.NewTodo(time.Now(), todoDesc)
		jsonData, err := json.Marshal(newTodo)
		todo.Add(jsonData, "tasks.json")
	case Update:
		todoId, err := c.id.Get()
		if err != nil {
			panic(err)
		}

		todo.Update(int(todoId), "tasks.json", c.desc, optional.String{})
	case InProcess:
		todoId, err := c.id.Get()
		if err != nil {
			panic(err)
		}

		todo.Update(int(todoId), "tasks.json", optional.String{}, optional.NewString(todo.StateToString(todo.InProcess)))
	case Done:
		todoId, err := c.id.Get()
		if err != nil {
			panic(err)
		}

		todo.MakeDone(int(todoId), "tasks.json")
	}
	return nil
}
