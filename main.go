package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <banner> <text>")
		return
	}

	bannerFile := os.Args[1]
	text := strings.Join(os.Args[2:], " ")

	lines, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	renderText(text, lines)
}
