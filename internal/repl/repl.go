package repl

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Woutjeee/go_dnd/internal/commands"
	"github.com/Woutjeee/go_dnd/internal/configuration"
)

func StartRepl(cfg *configuration.Config) {
	fmt.Println("Starting REPL session.")

	// Get all the current registered commands.
	emptyFlags := make(map[string]string)
	c := commands.GetCommands(cfg, emptyFlags)

	// Start a reader to get input of what the user entered.
	reader := bufio.NewScanner(os.Stdin)

	// Print the repl once.
	printRepl()

	// Start loop.
	for reader.Scan() {

		// Clean the input.
		text, flags := cleanInput(reader.Text())

		// Check if the command exists, if it does execute it.
		if commad, exists := c[text]; exists {
			cfg.LastCommand = commad
			cfg.LastFlags = flags
			err := commad.Command(cfg)
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			cfg.LastCommand = commad
		}
		// Print repl again.
		printRepl()
	}
}

var flagRegex = regexp.MustCompile(`(-\S+)\s*('[^']*'|\S+)?`)

func cleanInput(text string) (string, map[string]string) {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	parts := strings.Fields(output)
	commandName := parts[0]
	matches := flagRegex.FindAllStringSubmatch(output, -1)

	result := make(map[string]string)
	for _, match := range matches {
		flag := match[1]
		value := match[2]

		if value != "" {
			value = strings.Trim(value, "'")
			result[flag] = value
		}
	}

	return commandName, result
}

func printRepl() {
	fmt.Print("-> ")
}
