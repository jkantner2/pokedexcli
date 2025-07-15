package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		if len(cleanText) == 0 {
	
		}
		firstWord := cleanText[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}

func cleanInput(text string)[]string {
	words := strings.ToLower(text)
	output := strings.Fields(words)
	return output
}
