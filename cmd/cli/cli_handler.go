package cli

import (
	"fmt"
	"github.com/markphelps/optional"
	"time"
	"todo_cli/main/cmd/logger"
	"todo_cli/main/cmd/todo"
	"todo_cli/main/cmd/utils"
)

func (c *Config) Handle() error {
	dstFileName := "tasks.json"
	log, err := logger.Config()
	if err != nil {
		return err
	}

	switch c.command {
	case New:
		desc, err := c.desc.Get()
		if err != nil {
			panic(err)
		}

		id, err := todo.LastId(dstFileName)
		if err != nil {
			return err
		}

		newTodo := todo.NewTodo(id+1, time.Now(), desc)
		err = todo.Add(newTodo, dstFileName)
		if err != nil {
			return err
		}

		log.Info("Task successfully created")
		fmt.Println("Task successfully created")
	case Update:
		todoId, err := c.id.Get()
		if err != nil {
			return err
		}

		err = todo.Update(int(todoId), dstFileName, c.desc, optional.String{})
		if err != nil {
			return err
		}

		log.Info("Task successfully updated")
		fmt.Println("Task successfully created")

	case InProcess:
		todoId, err := c.id.Get()
		if err != nil {
			return err
		}

		err = todo.Update(
			int(todoId), dstFileName,
			optional.String{},
			optional.NewString(todo.InProcess.ToString()),
		)
		if err != nil {
			return err
		}

		log.Info("Task successfully updated")
		fmt.Println("Task successfully created")
	case Done:
		todoId, err := c.id.Get()
		if err != nil {
			return err
		}

		err = todo.MakeDone(int(todoId), dstFileName)
		if err != nil {
			return err
		}

		log.Info("Task completed successfully")
		fmt.Println("Task successfully created")
	case List:
		todos, err := todo.List(dstFileName)
		if err != nil {
			return err
		}

		for _, v := range todos {
			fmt.Println("Description:\t", v.Description)
			fmt.Println("ID:\t\t", v.ID)
			fmt.Println("Created at:\t", utils.ParseDate(v.CreatedAt))
			fmt.Println("Updated at:\t", utils.ParseDate(v.UpdatedAt))
			fmt.Println("State:\t\t", v.State)
			fmt.Println()
		}
	}
	log.Info("List of tasks is displayed")

	return nil
}
