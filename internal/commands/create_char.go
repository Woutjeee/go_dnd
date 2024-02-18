package commands

import (
	"fmt"

	"github.com/Woutjeee/go_dnd/internal/configuration"
)

func Create(cfg *configuration.Config, flags map[string]string) error {
	fmt.Println(flags)

	return nil
}
