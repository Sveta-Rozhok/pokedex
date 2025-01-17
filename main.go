package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Sveta-Rozhok/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient *pokeapi.Client
	nextURL       *string
	previousURL   *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := parts[1:]

		switch command {
		case "explore":
			if len(args) != 1 {
				fmt.Println("Usage: explore <location-area>")
				continue
			}
			commandExplore(&cfg, args[0])
		case "pokedex":
			commandPokedex(&cfg)
		case "catch":
			if len(args) != 1 {
				fmt.Println("Usage: catch <pokemon-name>")
				continue
			}
			commandCatch(&cfg, args[0])
		case "inspect":
			if len(args) != 1 {
				fmt.Println("Usage: inspect <pokemon-name>")
				continue
			}
			commandInspect(&cfg, args[0])
		case "map":
			commandMap(&cfg)
		case "mapb":
			commandMapb(&cfg)
		case "help":
			commandHelp()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Type 'help' for available commands")
		}
	}
}

func commandExplore(cfg *config, locationArea string) {
	fmt.Printf("Exploring %s...\n", locationArea)
	locationInfo, err := cfg.pokeapiClient.GetLocationArea(locationArea)
	if err != nil {
		fmt.Printf("Error exploring location: %v\n", err)
		return
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandRegistry = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
}

func commandMap(cfg *config) {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextURL)
	if err != nil {
		fmt.Println("Error fetching locations:", err)
		return
	}

	cfg.nextURL = resp.Next
	cfg.previousURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
}

func commandMapb(cfg *config) {
	if cfg.previousURL == nil {
		fmt.Println("You're on the first page")
		return
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousURL)
	if err != nil {
		fmt.Println("Error fetching locations:", err)
		return
	}

	cfg.nextURL = resp.Next
	cfg.previousURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
}

func commandCatch(cfg *config, pokemonName string) {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		fmt.Printf("Failed to get information about %s\n", pokemonName)
		return
	}

	catchRate := calculateCatchRate(pokemon.BaseExperience)

	if rand.Intn(100) <= catchRate {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
}

func calculateCatchRate(baseExp int) int {
	maxRate := 100
	rate := maxRate - (baseExp / 4)
	if rate < 5 {
		return 5
	}
	if rate > 90 {
		return 90
	}
	return rate
}

func commandInspect(cfg *config, pokemonName string) {
	pokemon, exists := cfg.caughtPokemon[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, ptype := range pokemon.Types {
		fmt.Printf("  - %s\n", ptype.Type.Name)
	}
}

func commandPokedex(cfg *config) {
	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemonName)
	}
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Print("Pokedex > ")
// 		scanner.Scan()
// 		input := scanner.Text()
// 		trimmedText := strings.TrimSpace(input)
// 		lowerText := strings.ToLower(trimmedText)

// 		if cmd, found := commandRegistry[lowerText]; found {
// 			if err := cmd.callback(); err != nil {
// 				fmt.Println("Error executing command:", err)
// 			}
// 		} else {
// 			fmt.Println("Unknown command")
// 		}
// 	}
// }
