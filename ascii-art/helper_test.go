// helper_test.go
package main

import "testing"

func loadStandard(t *testing.T) map[rune][]string {
	t.Helper()

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	return banner
}
