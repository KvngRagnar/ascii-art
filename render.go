package main

import "fmt"

func renderLine(text string, lines []string) {

	for row := 0; row < 8; row++ {

		for _, char := range text {
			ascii := int(char)
			index := ascii - 32
			start := index * 8

			fmt.Println(lines[start+row])

		}
		fmt.Println()
	}
}
