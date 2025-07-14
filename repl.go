package main

import (
	"strings"
)

func cleanInput(text string)[]string {
	words := strings.ToLower(text)
	output := strings.Fields(words)
	return output
}
