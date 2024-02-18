package commands

import (
	"os"

	"github.com/Woutjeee/go_dnd/internal/configuration"
)

func Exit(cfg *configuration.Config, flags map[string]string) error {
	os.Exit(0)
	return nil
}
