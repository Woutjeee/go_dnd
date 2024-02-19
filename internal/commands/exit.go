package commands

import (
	"os"

	"github.com/Woutjeee/go_dnd/internal/configuration"
)

func Exit(cfg *configuration.Config) error {
	os.Exit(0)
	return nil
}
