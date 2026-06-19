package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <text> <banner>")
		return
	}

	bannerFile := os.Args[2]
	text := os.Args[1]

	lines, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	renderText(text, lines)
}
