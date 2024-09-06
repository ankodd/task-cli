package cli

import (
	"errors"
	"github.com/markphelps/optional"
	"os"
	"strconv"
)

type Config struct {
	command Command
	id      optional.Int64
	desc    optional.String
}

func NewConfig() (Config, error) {
	args := os.Args
	if len(args) == 1 {
		return Config{}, errors.New("Missing args\n")
	}

	var desc optional.String
	var id optional.Int64

	command, err := CommandFromString(args[1])
	if err != nil {
		panic(err)
	}

	switch len(args) {
	case 2:
	case 3:
		if command == New {
			desc = optional.NewString(args[2])
			break
		}

		v, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			return Config{}, errors.New("Invalid id\n")
		}

		id = optional.NewInt64(v)
	case 4:
		v, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			return Config{}, errors.New("Invalid id\n")
		}

		id = optional.NewInt64(v)
		desc = optional.NewString(args[3])
	default:
		return Config{}, errors.New("Invalid number of arguments\n")
	}

	return Config{
		command: command,
		id:      id,
		desc:    desc,
	}, nil
}
