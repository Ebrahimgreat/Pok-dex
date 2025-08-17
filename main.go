package main

import (
	"bufio"
	"fmt"
	"strings"

	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}
func commandHelp() error {
	fmt.Println("Helo I can help you boss")
	return nil
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Helping user",
		callback:    commandHelp,
	},
}

func main() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the pokedex")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		if cmd, ok := commands[input]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Println("Error", err)
			}
		}
	}

}
