package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationApiResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

const locationEndpoint string = "https://pokeapi.co/api/v2/location/"

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
	locationUrl := locationEndpoint
	if cfg.Next != "" {
		locationUrl = cfg.Next
	}
	response, err := fetchLocation(locationUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	cfg.Next = response.Next
	cfg.Previous = response.Previous
	return nil
}

func commandPrevious(cfg *config) error {
	if cfg.Previous == "" {
		return errors.New("previous page not available")
	}
	response, err := fetchLocation(cfg.Previous)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	cfg.Next = response.Next
	cfg.Previous = response.Previous
	return nil
}

func fetchLocation(url string) (locationApiResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationApiResponse{}, errors.New("cannot fetch locations")
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		error := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		fmt.Println(error)
		return locationApiResponse{}, errors.New(error)
	}
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return locationApiResponse{}, err

	}
	responseJson := locationApiResponse{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		fmt.Println(err)
		return locationApiResponse{}, err
	}
	return responseJson, nil
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
		err := command.callback(&cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
