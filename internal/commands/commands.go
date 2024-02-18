package commands

import "github.com/Woutjeee/go_dnd/internal/configuration"

type Command struct {
	Name           string
	Description    string
	AvailableFlags map[string]string
	Command        func(cfg *configuration.Config) error
}

func GetCommands(cfg *configuration.Config) map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "A command that helps the user by printing the available commands.",
			AvailableFlags: map[string]string{
				"-h": "test flag",
			},
			Command: Help,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the app.",
			Command:     Exit,
		},
		"create character": {
			Name:        "create character",
			Description: "Starts the character creation proccess.",
			Command:     Create,
		},
	}
}
