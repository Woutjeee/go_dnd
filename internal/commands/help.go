package commands

import (
	"fmt"
)

func Help() error {
	fmt.Println("The following commands are available.")

	currCommands := GetCommands()

	for i := 0; i < len(currCommands); i++ {

	}

	return nil
}
