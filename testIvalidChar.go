package main

import "testing"

func TestInvalidChar(t *testing.T) {

	lines := mockBanner()

	block := getCharBlock(5, lines) // non-printable

	if len(block) != 8 {
		t.Error("invalid char should still return 8-line block")
	}
}
