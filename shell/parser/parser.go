package parser

import (
	"strings"
	"unicode"
)

func Parse(input string) [][]string {
	var commands [][]string
	var currentCommand []string
	var currentArg strings.Builder
	var inQuote rune

	for _, r := range input {
		switch {
		case r == inQuote:
			inQuote = 0
		case inQuote != 0:
			currentArg.WriteRune(r)
		case r == '|':
			// Split pipeline
			flushArg(&currentArg, &currentCommand)
			commands = append(commands, currentCommand)
			currentCommand = nil
		case unicode.IsSpace(r):
			flushArg(&currentArg, &currentCommand)
		case r == '"' || r == '\'':
			inQuote = r
		case r == '>' || r == '<':
			// Handle redirection
			flushArg(&currentArg, &currentCommand)
			currentCommand = append(currentCommand, string(r))
		default:
			currentArg.WriteRune(r)
		}
	}

	flushArg(&currentArg, &currentCommand)
	if len(currentCommand) > 0 {
		commands = append(commands, currentCommand)
	}
	return commands
}

func flushArg(b *strings.Builder, args *[]string) {
	if b.Len() > 0 {
		*args = append(*args, b.String())
		b.Reset()
	}
}