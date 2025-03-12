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

// TestOurByteBufferWrite tests that writing to a buffer works correctly
func TestOurByteBufferWrite(t *testing.T) {
	// Create a buffer with initial content
	initialContent := []byte("hello")
	buffer := NewOurByteBuffer(initialContent)

	// Write additional content to the buffer
	additionalContent := []byte(" world")
	n, err := buffer.Write(additionalContent)

	//check for unexpected errors
	if err != nil {
		t.Fatalf("Unexpected error writing to buffer: %v", err)
	}
	// verify the write function reported the correct number of bytes written
	if n != len(additionalContent) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(additionalContent), n)
	}

	// check if the buffer contains both the initial and additional content
	expected := []byte("hello world")
	result := buffer.Bytes()
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected buffer to contain %q, got %q", expected, result)
	}
}
