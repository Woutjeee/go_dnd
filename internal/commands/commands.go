package commands

import "github.com/Woutjeee/go_dnd/internal/configuration"

func GetCommands(cfg *configuration.Config) map[string]configuration.Command {
	return map[string]configuration.Command{
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
		"create": {
			Name:        "create character",
			Description: "Starts the character creation proccess.",
			AvailableFlags: map[string]string{
				"-n": "name of character",
			},
			Command: Create,
		},
	}
}
