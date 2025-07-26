package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
		callback:	nmap,
    		},
	    "mapb": {
		name:		"mapb",
		description:	"step backwards through list of map locations",
		callback:	mapb,
		},
    }
}

type config struct {
    pokeapiClient	pokeapi.Client
    nextURL		string
    prevURL		string
}

var currentConfig = config {
	nextURL:	"https://pokeapi.co/api/v2/location-area",
	prevURL:	"",
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



func nmap(cfg *config) error {
	res, err := http.Get(config.nextURL)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	var mapInformation mapInfo
	if err := json.Unmarshal(data, &mapInformation); err != nil {
		return fmt.Errorf("error unmarshaling json data: %w", err)
	}
	
	config.nextURL = mapInformation.Next
	config.prevURL = mapInformation.Previous

	var listOfNames []string

	for _, result := range mapInformation.Results{
		listOfNames = append(listOfNames, result.Name)
	}

	for _, name := range listOfNames {
		fmt.Println(name)
	}

	return nil
}

func mapb(cfg *config) error {
	res, err := http.Get(config.prevURL)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	var mapInformation mapInfo
	if err := json.Unmarshal(data, &mapInformation); err != nil {
		return fmt.Errorf("error unmarshaling json data: %w", err)
	}
	
	config.nextURL = mapInformation.Next
	config.prevURL = mapInformation.Previous

	var listOfNames []string

	for _, result := range mapInformation.Results{
		listOfNames = append(listOfNames, result.Name)
	}

	for _, name := range listOfNames {
		fmt.Println(name)
	}
	return nil
}
