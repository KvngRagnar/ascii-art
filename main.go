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

	// Debug prints
	fmt.Println("Total lines:", len(lines))
	fmt.Println("First line:", lines[0])
}
