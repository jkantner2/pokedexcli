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
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg)
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
    callback	func(config *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
	    "exit": {
    		name:		"exit",
		description:	"Exit the Pokedex",
		callback:	commandExit,
    		},

	    "help": {
		name:		"help",
		description:	"display manual page listing commands and their function",
		callback:	commandHelp,
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
    }
}

type config struct {
    pokeapiClient		pokeapi.Client
    nextLocationsURL		*string
    prevLocationsURL		*string
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

