package main

import (
	"testing"
)

// TestValidateInput_ValidStrings checks that fully valid inputs return 0, nil.
func TestValidateInput_ValidStrings(t *testing.T) {
	valid := []string{
		"Hello",
		"hello world",
		"Hello, World!",
		"1234567890",
		"~`!@#$%^&*()-_=+[]{}|;':\",./<>?",
		" ",
		"",
	}
	for _, s := range valid {
		r, err := ValidateInput(s)
		if err != nil {
			t.Errorf("ValidateInput(%q): unexpected error %v (rune %q)", s, err, r)
		}
		if r != 0 {
			t.Errorf("ValidateInput(%q): expected rune 0, got %q", s, r)
		}
	}
}

// TestValidateInput_InvalidCharacters checks that each non-ASCII or
// below-32 character is caught and returned.
func TestValidateInput_InvalidCharacters(t *testing.T) {
	cases := []struct {
		input       string
		invalidRune rune
	}{
		{"héllo", 'é'},
		{"naïve", 'ï'},
		{"café", 'é'},
		{"こんにちは", 'こ'},
		{"\x01hello", '\x01'},
		{"\x00", '\x00'},
		{"\x1F", '\x1F'},      // ASCII 31 — just below valid range
		{"hello\x7F", '\x7F'}, // ASCII 127 — just above valid range
		{"😀", '😀'},
	}
	for _, tc := range cases {
		r, err := ValidateInput(tc.input)
		if err == nil {
			t.Errorf("ValidateInput(%q): expected error, got nil", tc.input)
			continue
		}
		if r != tc.invalidRune {
			t.Errorf("ValidateInput(%q): expected rune %q, got %q", tc.input, tc.invalidRune, r)
		}
	}
}

// TestValidateInput_ReturnsFirstInvalidOnly checks that when there are
// multiple invalid characters, only the first one is returned.
func TestValidateInput_ReturnsFirstInvalidOnly(t *testing.T) {
	// 'é' comes before 'ñ' — must return 'é'
	r, err := ValidateInput("héñ")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if r != 'é' {
		t.Errorf("expected first invalid rune 'é', got %q", r)
	}
}

// TestValidateInput_BoundaryASCII32And126 checks that ASCII codes 32 (space)
// and 126 (~) are accepted — they are the inclusive edges of the valid range.
func TestValidateInput_BoundaryASCII32And126(t *testing.T) {
	boundaries := []rune{32, 126}
	for _, code := range boundaries {
		r, err := ValidateInput(string(code))
		if err != nil {
			t.Errorf("ASCII %d (%q) should be valid, got error: %v (rune %q)", code, code, err, r)
		}
	}
}

// TestValidateInput_BoundaryASCII31And127 checks that ASCII codes 31 and 127
// are rejected — they sit just outside the valid range.
func TestValidateInput_BoundaryASCII31And127(t *testing.T) {
	for _, code := range []rune{31, 127} {
		_, err := ValidateInput(string(code))
		if err == nil {
			t.Errorf("ASCII %d should be invalid, but got no error", code)
		}
	}
}

// TestValidateInput_ValidPrefixBeforeInvalid checks that even if most of the
// string is valid, the function still catches an invalid char at the end.
func TestValidateInput_ValidPrefixBeforeInvalid(t *testing.T) {
	_, err := ValidateInput("Hello World é")
	if err == nil {
		t.Error("expected error for string with invalid char at end, got nil")
	}
}

// TestValidateInput_EmptyString expects 0, nil for empty input.
func TestValidateInput_EmptyString(t *testing.T) {
	r, err := ValidateInput("")
	if err != nil {
		t.Errorf("ValidateInput(\"\"): unexpected error %v", err)
	}
	if r != 0 {
		t.Errorf("ValidateInput(\"\"): expected rune 0, got %q", r)
	}
}
