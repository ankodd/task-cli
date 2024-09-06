package cli

import "errors"

type Command string

const (
	New       Command = "new"
	Update    Command = "update"
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
	case "list":
		return List, nil
	default:
		return "", errors.New("Invalid command\n")
	}
}
