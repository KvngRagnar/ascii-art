package main

import "testing"

func TestLoadBannerValid(t *testing.T) {
	lines, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected banner to load successfully, got %v", err)
	}

	if len(lines) < 855 {
		t.Fatalf("expected at least 855 lines, got %d", len(lines))
	}
}

func TestLoadBannerFileDoesNotExist(t *testing.T) {
	_, err := loadBanner("does-not-exist.txt")

	if err == nil {
		t.Fatal("expected an error for a non-existent banner file")
	}
}

func TestLoadBannerEmptyFile(t *testing.T) {
	_, err := loadBanner("empty.txt")

	if err == nil {
		t.Fatal("expected an error for an empty banner file")
	}
}

func TestGetCharBlockPrintableCharacter(t *testing.T) {
	lines, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	block := getCharBlock('A', lines)

	if len(block) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(block))
	}
}

func TestGetCharBlockSpace(t *testing.T) {
	lines, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	block := getCharBlock(' ', lines)

	if len(block) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(block))
	}
}

func TestGetCharBlockInvalidCharacter(t *testing.T) {
	lines, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	block := getCharBlock('\t', lines)

	if len(block) != 8 {
		t.Fatalf("expected 8 blank lines, got %d", len(block))
	}

	for _, line := range block {
		if line != "        " {
			t.Fatalf("expected blank line, got %q", line)
		}
	}
}

func TestRenderLine(t *testing.T) {
	lines, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Test passes if no panic occurs.
	renderLine("Hello", lines)
}

func TestRenderText(t *testing.T) {
	lines, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Test passes if no panic occurs.
	renderText("Hello", lines)
}

func TestRenderTextWithNewLine(t *testing.T) {
	lines, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Test passes if no panic occurs.
	renderText("Hello\\nWorld", lines)
}

func TestRenderTextOnlyNewLine(t *testing.T) {
	lines, err := loadBanner("standard.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Test passes if no panic occurs.
	renderText("\\n", lines)
}
