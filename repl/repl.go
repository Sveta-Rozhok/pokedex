package repl

import (
	"strings"
)

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	lowerText := strings.ToLower(trimmedText)
	words := strings.Fields(lowerText)

	return words
}
