package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the banner file
	data, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	// Convert bytes to string
	content := string(data)

	// Split into lines
	lines := strings.Split(content, "\n")
	text := "HI"

	for row := 0; row < 8; row++ {
		for _, char := range text {
			index := int(char) - 32
			start := index * 8
			fmt.Print(lines[start+row])
		}
		fmt.Println()
	}
}
