package namer

import (
	"regexp"
	"testing"
)

func TestGenerateName(t *testing.T) {
	const prefix = "vistar"

	result := GenerateName()

	pattern := `^vistar_\d{8}_\d{6}$`
	matched, err := regexp.MatchString(pattern, result)
	if err != nil {
		t.Fatalf("Failed to compile regex: %v", err)
	}

	if !matched {
		t.Errorf("Generated name '%s' does not match expected pattern '%s'", result, pattern)
	}

	if len(result) < len(prefix) || result[:len(prefix)] != prefix {
		t.Errorf("Generated name '%s' does not start with prefix '%s'", result, prefix)
	}
}
