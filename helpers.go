package main

import (
	"os"
	"strings"
)

// Check if the value is empty
func IsSettingEmpty(ev string) string {
	s := os.Getenv(ev)
	// Fail if there is no value for that environment variable
	if !DoesValueExists(s) {
		Fatalf("Mandatory parameter \"%s\" is missing from the environment variable or its empty", ev)
	}
	return s
}

// Trim spaces and provide the value
func DoesValueExists(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}
	return true
}

