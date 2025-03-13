package main

import (
	"bytes"
	"testing"
)

func TestFilteringPipe(t *testing.T) {
	// Create a buffer to capture the output
	output := &bytes.Buffer{}

	// Create our filtering pipe that will write to the buffer
	filter := NewFilteringPipe(output)

	// Write a string with digits to the filter
	input := []byte("abc123def456")
	n, err := filter.Write(input)

	// check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify the write method returned the original input length
	if n != len(input) {
		t.Errorf("Expected Write to return length %d, got %d", len(input), n)
	}

	// Check if the output has all digits removed
	expected := "abcdef"
	if output.String() != expected {
		t.Errorf("Expected output %q, got %q", expected, output.String())
	}
}
