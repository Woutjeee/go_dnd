package commmands

import (
	"fmt"
	"os"
)

type Command struct {
	Name    string
	Command func() error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:    "help",
			Command: help,
		},
		"exit": {
			Name:    "exit",
			Command: exit,
		},
	}

}

func exit() error {
	fmt.Println("Stopping REPL session.")
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Println("Help command executed")

	return nil
}
