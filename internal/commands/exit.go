package commands

import (
	"os"
)

func exit() error {
	os.Exit(0)
	return nil
}
