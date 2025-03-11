package main

import (
	"bytes"
	"testing"
)

// TestOurByteBufferInit verifies a new buffer contains the correct initial data
func TestOurByteBufferInit(t *testing.T) {
	// Create test data
	initialContent := []byte("hello world")
	// create a new buffer with our test data
	buffer := NewOurByteBuffer(initialContent)

	// Get the buffer's content
	result := buffer.Bytes()

	// Check if the content matches what we put in
	if !bytes.Equal(result, initialContent) {
		t.Errorf("Expected buffer to contain %q, got %q", initialContent, result)
	}
}
