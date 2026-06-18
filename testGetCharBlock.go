package main

import "testing"

func TestGetCharBlock(t *testing.T) {

	lines := mockBanner()

	block := getCharBlock('A', lines)

	if len(block) != 8 {
		t.Error("expected 8-line block")
	}
}
