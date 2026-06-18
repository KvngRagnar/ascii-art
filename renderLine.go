package main

import "fmt"

func renderLine(text string, lines []string) {

	for row := 0; row < 8; row++ {

		for _, char := range text {

			block := getCharBlock(char, lines)

			fmt.Print(block[row])
		}

		fmt.Println()
	}
}
