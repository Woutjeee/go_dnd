package commands

import (
	"fmt"
)

func Help() error {
	fmt.Println("The following commands are available.")

	for _, command := range GetCommands() {
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
