package commands

import (
	"bufio"
	"fmt"
	"os"
)

func Create() error {
	fmt.Println("Enter characters name")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	return nil
}
