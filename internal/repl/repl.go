package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Woutjeee/go_dnd/internal/commmands"
)

func StartRepl() {
	fmt.Println("Starting REPL session.")

	// Get all the current registered commands.
	c := commmands.GetCommands()

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
		} else {
			// Command did not exists.
			fmt.Println("Unknown command")
		}

		// Print repl again.
		printRepl()
	}
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func printRepl() {
	fmt.Print("-> ")
}