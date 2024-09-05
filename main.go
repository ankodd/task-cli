package main

import (
	"fmt"
	"todo_cli/main/cli"
)

func main() {
	config, err := cli.NewConfig()
	if err != nil {
		fmt.Printf(err.Error())
	}

	err = config.Handle()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
