package main

import (
	"os"
	"strings"
)

func loadBanner(filename string) ([]string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	return lines, nil
}
