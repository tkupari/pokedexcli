package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {

	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}

func commandExit() error {
	fmt.Println("Exiting program..")
	os.Exit(0)
	return nil
}

func main() {
	commands := getCommands()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex > ")
		reader.Scan()
		input := reader.Text()

		command, ok := commands[input]
		if !ok {
			fmt.Println("invalid command: " + input)
			continue
		}
		command.callback()
	}
}
