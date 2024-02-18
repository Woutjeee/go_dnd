package commands

import (
	"fmt"
	"os"
)

func exit() error {
	fmt.Println("Stopping REPL session.")
	os.Exit(0)
	return nil
}
