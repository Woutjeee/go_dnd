package main

import (
	"github.com/Woutjeee/go_dnd/internal/configuration"
	"github.com/Woutjeee/go_dnd/internal/repl"
)

func main() {
	config := configuration.Config{
		LastCommand: "none",
	}

	repl.StartRepl(&config)
}
