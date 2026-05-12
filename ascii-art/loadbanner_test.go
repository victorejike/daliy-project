package main

import (
	"os"

	"strings"

	"testing"
)

// TestLoadBanner_ReturnsCorrectCharCount checks that all 95 printable ASCII

// characters (codes 32–126) are present in the loaded banner map.

func TestLoadBanner_ReturnsCorrectCharCount(t *testing.T) {

	banner, err := LoadBanner("standard.txt")

	if err != nil {

		t.Fatalf("unexpected error loading standard.txt: %v", err)

	}

	if len(banner) != 95 {

		t.Errorf("expected 95 characters in banner, got %d", len(banner))

	}

}

// TestLoadBanner_EachCharHasEightLines checks that every character in the

// banner map has exactly 8 art lines.

func TestLoadBanner_EachCharHasEightLines(t *testing.T) {

	banner, err := LoadBanner("standard.txt")

	if err != nil {

		t.Fatalf("unexpected error loading standard.txt: %v", err)

	}

	for r, art := range banner {

		if len(art) != 8 {

			t.Errorf("character %q: expected 8 lines, got %d", r, len(art))

		}

	}

}

// TestLoadBanner_AllPrintableASCIIPresent checks every rune from 32 to 126

// individually to make sure none are missing.

func TestLoadBanner_AllPrintableASCIIPresent(t *testing.T) {

	banner, err := LoadBanner("standard.txt")

	if err != nil {

		t.Fatalf("unexpected error loading standard.txt: %v", err)

	}

	for code := rune(32); code <= 126; code++ {

		if _, ok := banner[code]; !ok {

			t.Errorf("character %q (ASCII %d) is missing from banner", code, code)

		}

	}

}

// TestLoadBanner_SpaceCharPresent specifically checks the space character

// since it is invisible and students often accidentally skip it.

func TestLoadBanner_SpaceCharPresent(t *testing.T) {

	banner, err := LoadBanner("standard.txt")

	if err != nil {

		t.Fatalf("unexpected error loading standard.txt: %v", err)

	}

	art, ok := banner[' ']

	if !ok {

		t.Fatal("space character (ASCII 32) is missing from banner")

	}

	if len(art) != 8 {

		t.Errorf("space character: expected 8 lines, got %d", len(art))

	}

}

// TestLoadBanner_NoLineContainsNewline checks that none of the stored art

// lines themselves contain a newline byte — the parser must strip them.

func TestLoadBanner_NoLineContainsNewline(t *testing.T) {

	banner, err := LoadBanner("standard.txt")

	if err != nil {

		t.Fatalf("unexpected error loading standard.txt: %v", err)

	}

	for r, art := range banner {

		for i, line := range art {

			if strings.Contains(line, "\n") {

				t.Errorf("character %q line %d contains a newline byte", r, i)

			}

		}

	}

}

// TestLoadBanner_KnownCharacters spot-checks a handful of characters that

// have well-known first lines in the standard banner.

func TestLoadBanner_KnownCharacters(t *testing.T) {

	banner, err := LoadBanner("standard.txt")

	if err != nil {

		t.Fatalf("unexpected error loading standard.txt: %v", err)

	}

	cases := []struct {
		char rune

		lineIdx int

		contains string
	}{

		{'A', 0, "_"}, // top of 'A' has an underscore

		{'a', 0, ""}, // first line of lowercase a is blank in standard

		{'Z', 0, "_"},

		{'!', 0, "_"},
	}

	for _, tc := range cases {

		art, ok := banner[tc.char]

		if !ok {

			t.Errorf("character %q not found in banner", tc.char)

			continue

		}

		if len(art) <= tc.lineIdx {

			t.Errorf("character %q: not enough lines (need index %d)", tc.char, tc.lineIdx)

			continue

		}

		if tc.contains != "" && !strings.Contains(art[tc.lineIdx], tc.contains) {

			t.Errorf("character %q line %d: expected to contain %q, got %q",

				tc.char, tc.lineIdx, tc.contains, art[tc.lineIdx])

		}

	}

}

// TestLoadBanner_FileNotFound expects a non-nil error and a nil map when the

// file does not exist.

func TestLoadBanner_FileNotFound(t *testing.T) {

	banner, err := LoadBanner("notfound.txt")

	if err == nil {

		t.Error("expected a non-nil error for missing file, got nil")

	}

	if banner != nil {

		t.Error("expected nil map on error, got non-nil")

	}

}

// TestLoadBanner_EmptyFile expects an error when the file exists but is empty.

func TestLoadBanner_EmptyFile(t *testing.T) {

	f, err := os.CreateTemp("", "empty_banner_*.txt")

	if err != nil {

		t.Fatal(err)

	}

	defer os.Remove(f.Name())

	f.Close()

	_, err = LoadBanner(f.Name())

	if err == nil {

		t.Error("expected error for empty file, got nil")

	}

}

// TestLoadBanner_InvalidContent expects an error when the file has random

// content that does not conform to the 8-lines-per-character format.

func TestLoadBanner_InvalidContent(t *testing.T) {

	f, err := os.CreateTemp("", "invalid_banner_*.txt")

	if err != nil {

		t.Fatal(err)

	}

	defer os.Remove(f.Name())

	_, _ = f.WriteString("this is not a valid banner file\nonly two lines\n")

	f.Close()

	_, err = LoadBanner(f.Name())

	if err == nil {

		t.Error("expected error for invalid banner content, got nil")

	}

}

// TestLoadBanner_ShadowAndThinkertoy verifies the other two banner files also

// load correctly with 95 chars each having 8 lines.

func TestLoadBanner_ShadowAndThinkertoy(t *testing.T) {

	for _, filename := range []string{"shadow.txt", "thinkertoy.txt"} {

		banner, err := LoadBanner(filename)

		if err != nil {

			t.Errorf("%s: unexpected error: %v", filename, err)

			continue

		}

		if len(banner) != 95 {

			t.Errorf("%s: expected 95 chars, got %d", filename, len(banner))

		}

		for r, art := range banner {

			if len(art) != 8 {

				t.Errorf("%s: character %q has %d lines, expected 8", filename, r, len(art))
			}
		}
	}
}
