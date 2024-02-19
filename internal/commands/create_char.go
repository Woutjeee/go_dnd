package commands

import (
	"fmt"

	"github.com/Woutjeee/go_dnd/internal/configuration"
)

func Create(cfg *configuration.Config) error {
	fmt.Println(cfg.LastFlags)

	return nil
}
