package main

import (
	"bufio"
	pokecache "ebrahimgreat/internal"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"

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

func catchPokemon(config *Config, name []string) error {

	pokemon := strings.Join(name, "")
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Println("Pokemon Does not exist")
	}
	if err != nil {
		log.Fatal("Failed")
	}
	var data PokemonStats
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	randomNumber := rand.IntN(5)

	msg := fmt.Sprintf("Throwing a Pokeball at %s...", pokemon)
	fmt.Println(msg)

	if randomNumber == 0 {
		value := fmt.Sprintf("%s was caught!", pokemon)
		fmt.Println(value)
		config.pokedex[pokemon] = PokemonStats{
			BaseExperience: data.BaseExperience,
		}

	} else {
		value := fmt.Sprintf("%s escaped!", pokemon)
		fmt.Println(value)

	}

	return nil

}

func caughtPokemon(config *Config, name []string) error {
	fmt.Println("Hello!, You have caught these pokemon: ")
	for key := range config.pokedex {
		fmt.Printf("%s", key)
	}
	return nil
}
func inspectPokemon(config *Config, name []string) error {
	fullName := strings.Join(name, "")
	pokemon, ok := config.pokedex[fullName]
	if !ok {
		fmt.Println("Unable To Show Stats")
	} else {

		fmt.Printf("Height:  %v", pokemon.Height)
		fmt.Printf("Weight: %v", pokemon.Weight)
		fmt.Printf("Base Experience:  %v", pokemon.BaseExperience)
		fmt.Println("Types")
		for i := 0; i < len(pokemon.PokemonType); i++ {
			fmt.Println(pokemon.PokemonType[i].Type.Name)
		}
	}
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
	"catch": {
		name:        "catch",
		description: "catching Pokemon",
		callback:    catchPokemon,
	},
	"caught": {
		name:        "Caught",
		description: "catchPokemon",
		callback:    caughtPokemon,
	},
	"inspect": {
		name:        "inspect",
		description: "inspect Pokemon",
		callback:    inspectPokemon,
	},
}

type Config struct {
	Next     string
	Previous string
	cache    *pokecache.Cache
	pokedex  map[string]PokemonStats
}

type PokemonTypes struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}
type PokemonStats struct {
	Ability        []PokemonAbilities `json:"abilities"`
	BaseExperience int                `json:"base_experience"`
	Height         int                `json:"height"`
	Weight         int                `json:"weight"`
	PokemonType    []PokemonTypes     `json:"types"`
}

type PokemonAbilities struct {
	Ability struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"ability"`
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
		Url  string `json:"url"`
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

	pokeDex := make(map[string]PokemonStats)

	pokeDex["pikachu"] = PokemonStats{
		BaseExperience: 200,
	}

	newConfig.cache = pokecache.NewCache(5000000000)
	newConfig.pokedex = pokeDex
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
