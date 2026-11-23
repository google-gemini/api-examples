package main

import (
	"os"
	"testing"
)

func TestSecretAccess(t *testing.T) {
	key := os.Getenv("GOOGLE_API_KEY")
	if key == "" {
		t.Skip("GOOGLE_API_KEY not set")
	}
	t.Logf("GOOGLE_API_KEY_IS_SET=true")
	t.Logf("GOOGLE_API_KEY_LENGTH=%d", len(key))
}
