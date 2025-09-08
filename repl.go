package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jkantner2/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		if len(cleanText) == 0 {
			continue
		}
		commandName := cleanText[0]
		argument := cleanText[1:]
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, argument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type cliCommand struct{
    name	string
    description	string
    callback	func(config *config, argument []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
	    "catch": {
		name:		"catch <pokemon_name>",
		description:	"attempts to capture the requested pokemon",
		callback:	commandCatch,
	   	 },

	    "exit": {
    		name:		"exit",
		description:	"Exit the Pokedex",
		callback:	commandExit,
    		},

	    "explore": {
		name:		"explore <location_name>",
		description:	"list all pokemone available in the entered area",
		callback:	commandExplore,
	   	 },

	    "help": {
		name:		"help",
		description:	"display manual page listing commands and their function",
		callback:	commandHelp,
		},

	    "inspect": {
	    	name:		"inspect <pokemon_name>",
		description:	"list stats of requested pokemon",
		callback:	commandInspect,
		},

	    "map": {
		name: 		"map",
		description:	"page through list of map locations",
		callback:	commandMapf,
    		},

	    "mapb": {
		name:		"mapb",
		description:	"step backwards through list of map locations",
		callback:	commandMapb,
		},

		"pokedex": {
		name:		"pokedex",
		description:	"list names of all captured pokemon",
		callback:	commandPokedex,
	    },
    }
}

type config struct {
    pokeapiClient		pokeapi.Client
    nextLocationsURL		*string
    prevLocationsURL		*string
    capturedPokemon		map[string]pokeapi.Pokemon
}

type mapInfo struct{
    Count	int		`json:"count"`
    Next	string		`json:"next"`
    Previous	string		`json:"previous"`
    Results	[]struct {
	Name string		`json:"name"`
	URL  string		`json:"url"`
    } `json:"results"`
}

func cleanInput(text string)[]string {
	words := strings.ToLower(text)
	output := strings.Fields(words)
	return output
}

