package main

import (
	"fmt"
	"os"
	"todo_cli/main/cmd/cli"
	"todo_cli/main/cmd/logger"
)

func main() {
	log, err := logger.Config()
	if err != nil {
		fmt.Println("Error creating logger: ", err)
		os.Exit(1)
	}

	config, err := cli.NewConfig()
	if err != nil {
		log.Error(err.Error())
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = config.Handle()
	if err != nil {
		log.Error(err.Error())
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
