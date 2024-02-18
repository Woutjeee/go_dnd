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
	c := commands.GetCommands()

	// Start a reader to get input of what the user entered.
	reader := bufio.NewScanner(os.Stdin)

	// Print the repl once.
	printRepl()

	// Start loop.
	for reader.Scan() {

		// Clean the input.
		text := cleanInput(reader.Text())

		// Check if the command exists, if it does execute it.
		if commad, exists := c[text]; exists {
			err := commad.Command()
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		}
		// Print repl again.
		printRepl()
	}
}

var flagRegex = regexp.MustCompile(`(-\S+)\s*(".*?"|\S+)?`)

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	matches := flagRegex.FindAllStringSubmatch(output, -1)

	result := make(map[string]string)

	for _, match := range matches {
		flag := match[1]
		value := match[2]

		if value != "" {
			result[flag] = value
		}
	}

	fmt.Println(result)

	return output
}

func printRepl() {
	fmt.Print("-> ")
}
