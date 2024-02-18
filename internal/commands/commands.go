package commands

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Description string
	Command     func() error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "A command that helps the user by printing the available commands.",
			Command:     Help,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the app.",
			Command:     exit,
		},
		"create character": {
			Name:        "create character",
			Description: "Starts the character creation proccess.",
			Command:     Create,
		},
	}

}

func exit() error {
	fmt.Println("Stopping REPL session.")
	os.Exit(0)
	return nil
}
