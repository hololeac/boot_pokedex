package main

import (
	"strings"
)

func cleanInput(text string) []string {
	loweredInput := strings.ToLower(text)
	words := strings.Fields(loweredInput)

	return words
}
