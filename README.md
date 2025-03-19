# Pokedex CLI

Pokedex CLI is a command-line application written in Go that allows interaction with PokeAPI for exploring locations, catching Pokémon, and viewing information about them.

## Installation and Setup

### 1. Clone the repository
```sh
git clone https://github.com/Sveta-Rozhok/pokedex.git
cd pokedex
```

### 2. Install dependencies
```sh
go mod tidy
```

### 3. Build and run
```sh
go run main.go
```

## Commands

| Command      | Description |
|--------------|------------|
| `help`       | Displays the help for commands |
| `exit`       | Exits the Pokedex |
| `explore <location-area>` | Explores the specified location and lists the Pokémon found there |
| `catch <pokemon-name>` | Attempts to catch the specified Pokémon |
| `inspect <pokemon-name>` | Displays the stats of the caught Pokémon |
| `pokedex`    | Lists all caught Pokémon |
| `map`        | Displays the available locations |
| `mapb`       | Displays the previous list of locations |

## Example Usage
```sh
Pokedex > explore forest
Exploring forest...
Found Pokémon:
 - Pikachu
 - Bulbasaur

Pokedex > catch Pikachu
Throwing a Pokeball at Pikachu...
Pikachu was caught!

Pokedex > inspect Pikachu
Name: Pikachu
Height: 4
Weight: 60
Stats:
  - Speed: 90
  - Attack: 55
Types:
  - Electric

Pokedex > exit
Closing the Pokedex... Goodbye!
```

## Code Overview

Main components:
- `config` - Configuration structure containing the PokeAPI client and caught Pokémon data.
- `pokeapi.Client` - Client for interacting with the PokeAPI.
- `commandExplore` - Command for exploring locations.
- `commandCatch` - Command for catching Pokémon.
- `commandInspect` - Command for viewing the stats of caught Pokémon.
- `commandPokedex` - Command for displaying all caught Pokémon.
- `commandMap` and `commandMapb` - Commands for navigation through locations.

## Requirements
- Go 1.20+
- An internet connection to interact with PokeAPI

## Authors
Developed by [Sveta-Rozhok](https://github.com/Sveta-Rozhok)

## License
This project is licensed under the MIT License.

## 🤝 Contributing

### Clone the repo

```bash
git clone https://github.com/Sveta-Rozhok/pokedex.git
cd pokedex
```

### Build the project

```bash
go build
```

### Run the project

```bash
go run main.go
```

### Run the tests

```bash
go test ./...
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
