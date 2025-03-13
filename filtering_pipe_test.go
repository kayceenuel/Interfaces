package main

import (
	"bytes"
	"testing"
)

// TestFilteringPipe tests basic filtering functionality
func TestFilteringPipe(t *testing.T) {
	// Create a buffer to capture the output
	output := &bytes.Buffer{}

	// Create our filtering pipe that will write to the buffer
	filter := NewFilteringPipe(output)

	// Write a string with digits to the filter
	input := []byte("abc123def456")
	n, err := filter.Write(input)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify the Write method returned the original input length
	if n != len(input) {
		t.Errorf("Expected Write to return length %d, got %d", len(input), n)
	}

	// Check if the output has all digits removed
	expected := "abcdef"
	if output.String() != expected {
		t.Errorf("Expected output %q, got %q", expected, output.String())
	}
}

// TestFilteringPipeTable uses table-driven tests for multiple scenarios
func TestFilteringPipeTable(t *testing.T) {
	// Define test cases with inputs and expected outputs
	tests := []struct {
		name     string // Name of the test case
		input    string // Input string to filter
		expected string // Expected output after filtering
	}{
		{
			name:     "No digits",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:     "Only digits",
			input:    "12345",
			expected: "",
		},
		{
			name:     "Mixed content",
			input:    "abc123def456",
			expected: "abcdef",
		},
		{
			name:     "Digits with spaces",
			input:    "phone: 555-1234",
			expected: "phone: -",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Start and end with digits",
			input:    "123abc456",
			expected: "abc",
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			output := &bytes.Buffer{}

			// Create a filtering pipe writing to the buffer
			filter := NewFilteringPipe(output)

			// Write the test input to the filter
			_, err := filter.Write([]byte(tt.input))

			// Check for unexpected errors
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check if the output matches expectations
			if output.String() != tt.expected {
				t.Errorf("Expected output %q, got %q", tt.expected, output.String())
			}
		})
	}
}
