package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next string
	Previous string
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
		"map": {
			name:        "map",
			description: "List next 20 locations",
			callback:    commandNext,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 locations",
			callback:    commandPrevious,
		},
	}
}

func commandHelp(cfg *config) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Exiting program..")
	os.Exit(0)
	return nil
}

func commandNext(cfg *config) error {
	return nil
}

func commandPrevious(cfg *config) error {
	return nil
}

func main() {
	cfg := config{}
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
		command.callback(&cfg)
	}
}
