package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Woutjeee/go_dnd/internal/configuration"
)

func Create(cfg *configuration.Config) error {
	fmt.Println("Enter characters name")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	return nil
}
