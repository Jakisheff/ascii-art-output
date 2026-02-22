package main

import (
	"testing"
)

func TestVerifyBannerHash(t *testing.T) {
	err := VerifyBannerHash("standard")
	if err != nil {
		t.Errorf("Validation for standard banner failed: %v", err)
	}
}

func TestReadBanner(t *testing.T) {
	banner, err := ReadBanner("standard")
	if err != nil {
		t.Fatalf("Failed to read standard banner: %v", err)
	}

	if len(banner['A']) != 8 {
		t.Errorf("Expected 8 lines for character 'A', got %d", len(banner['A']))
	}
}

func TestConvertToASCII(t *testing.T) {
	banner, err := ReadBanner("standard")
	if err != nil {
		t.Fatalf("Failed to read standard banner: %v", err)
	}

	result := ConvertToASCII("H", banner)
	if result == "" {
		t.Errorf("Expected non-empty result for 'H'")
	}
}

func TestConvertToASCIINewline(t *testing.T) {
	banner, _ := ReadBanner("standard")

	result := ConvertToASCII("\\n", banner)
	if result != "\n" {
		t.Errorf("Expected strictly single newline, got %d length", len(result))
	}
}
