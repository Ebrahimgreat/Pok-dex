package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}

func commandHelp(config *Config) error {
	fmt.Println("Helo I can help you boss")
	return nil
}

func commandPrevious(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Previous != "" {
		url = config.Previous
	}
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response Failed")
	}
	if err != nil {
		log.Fatal(err)
	}

	var myData Data
	err = json.Unmarshal(body, &myData)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(myData.Results); i++ {
		fmt.Println(myData.Results[i].Name)
	}
	config.Next = myData.Next
	config.Previous = myData.Previous
	return nil
}

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != "" {
		url = config.Next
	}
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response Failed")
	}
	if err != nil {
		log.Fatal(err)
	}

	var myData Data
	err = json.Unmarshal(body, &myData)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(myData.Results); i++ {
		fmt.Println(myData.Results[i].Name)
	}
	config.Next = myData.Next
	config.Previous = myData.Previous
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
	"map": {
		name:        "map",
		description: "Map",
		callback:    commandMap,
	},
	"map-b": {
		name:        "map-b",
		description: "previous",
		callback:    commandPrevious,
	},
}

type Config struct {
	Next     string
	Previous string
}
type Data struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Welcome To Pokedex")
	fmt.Println("To See locations type: map")
	fmt.Println("To Go Back type: map-b")

	var newConfig Config
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		if cmd, ok := commands[input]; ok {
			cmd.callback(&newConfig)
		} else {
			fmt.Println("Error")
		}

	}

}
