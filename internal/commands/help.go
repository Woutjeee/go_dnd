package commands

import (
	"fmt"

	"github.com/Woutjeee/go_dnd/internal/configuration"
)

func Help(cfg *configuration.Config) error {
	fmt.Println("The following commands are available.")

	for _, command := range GetCommands(cfg) {
		fmt.Printf(`
--------------  %s  ----------------
The command name: %s
The command description: %s
------------------------------------
`, command.Name, command.Name, command.Description)
	}

	fmt.Println("To use the commands, enter the name of it and it will execute it.")
	return nil
}
