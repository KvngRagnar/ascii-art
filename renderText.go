package main

import "strings"

func renderText(text string, lines []string) {

	parts := strings.Split(text, "\\n")

	for _, part := range parts {
		renderLine(part, lines)
	}
}
