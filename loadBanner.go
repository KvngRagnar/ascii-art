package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("banner file is empty")
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 855 {
		return nil, fmt.Errorf("invalid banner file: expected 855 lines got %d", len(lines))
	}
	return lines, nil
}
