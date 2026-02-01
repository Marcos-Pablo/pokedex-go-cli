package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
