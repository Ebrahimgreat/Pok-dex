package main

import (
	"bufio"
	pokecache "ebrahimgreat/internal"
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
	callback    func(*Config, []string) error
}

func commandExit(config *Config, name []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}

func commandHelp(config *Config, name []string) error {
	fmt.Println("Helo I can help you boss")
	return nil
}

func commandPrevious(config *Config, name []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	var body []byte
	var err error

	if config.Previous != "" {
		url = config.Previous
	}
	cachedData, ok := config.cache.Get(url)
	if ok {
		fmt.Println("Printing Cached")
		body = cachedData
	} else {
		fmt.Println("Printing uncached")
		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response Failed")
		}
		if err != nil {
			log.Fatal(err)
		}
		config.cache.Add(url, body)
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

func explorePokemon(config *Config, name []string) error {

	location := strings.Join(name, "")
	if location == "" {
		fmt.Printf("Please enter a location")

	}
	fmt.Println(location)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	var body []byte
	var err error

	cachedData, ok := config.cache.Get(url)
	if ok {
		body = cachedData

	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)

		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			fmt.Println("Location Does not exist. Programme Ending.....")
		}
		if err != nil {
			log.Fatal("failed")
		}
		config.cache.Add(url, body)

	}
	var myData pokemonData
	err = json.Unmarshal(body, &myData)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(myData.Pokemon); i++ {
		fmt.Println(myData.Pokemon[i].Pokemon.Name)
	}
	return nil

}

func commandMap(config *Config, name []string) error {

	url := "https://pokeapi.co/api/v2/location-area"
	var body []byte
	var err error

	if config.Next != "" {
		url = config.Next
	}
	cachedData, ok := config.cache.Get(url)
	if ok {
		fmt.Println("Printing from Cache")
		body = cachedData

	} else {
		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response Failed")
		}
		if err != nil {
			log.Fatal(err)
		}
		config.cache.Add(url, body)
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
	"explore": {
		name:        "pokemon",
		description: "Calling Pokemon",
		callback:    explorePokemon,
	},
}

type Config struct {
	Next     string
	Previous string
	cache    *pokecache.Cache
}

type pokemonData struct {
	Pokemon []PokemonEncounter `json:"pokemon_encounters"`
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
type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		url  string `json:"url"`
	} `json:"pokemon"`
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
	newConfig.cache = pokecache.NewCache(5000000000)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		fields := strings.Fields(input)
		cmdName := fields[0]
		args := fields[1:]
		if cmd, ok := commands[cmdName]; ok {

			cmd.callback(&newConfig, args)
		} else {
			fmt.Println("Error")
		}

	}

}
