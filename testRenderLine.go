package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRenderLine(t *testing.T) {

	lines := mockBanner()

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	renderLine("A", lines)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	output := buf.String()

	if output == "" {
		t.Error("expected output but got empty string")
	}
}
