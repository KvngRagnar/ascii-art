package main

import (
	"fmt"
	"strings"
)

func renderText(text string, lines []string) {

	parts := strings.Split(text, "\\n")

	for i, part := range parts {
		if part == "" {
			if i == len(parts)-1 {
				continue
			}
			fmt.Println()
			continue

		}

		renderLine(part, lines)
	}
}
