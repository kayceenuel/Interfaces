package main

import (
	"bytes"
	"io"
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

func TestOurByteBufferReadAll(t *testing.T) {
	//create a buffer with test content
	initialContent := []byte("hello world")
	buffer := NewOurByteBuffer(initialContent)

	//create a destination slice large enough to hold all content
	readBuffer := make([]byte, len(initialContent))
	//Read from our buffer into the destination slice
	n, err := buffer.Read(readBuffer)

	//check for unexpected errors (EOF is expected after reading everything)
	if err != nil && err != io.EOF {
		t.Fatalf("Unexpected error reading from buffer: %v", err)
	}

	// verify we read the expected number of bytes
	if n != len(initialContent) {
		t.Errorf("Expected to read %d bytes, read %d", len(initialContent), n)
	}

	//check if what we read matches the original content
	if !bytes.Equal(readBuffer, initialContent) {
		t.Errorf("Expected to read %q, got %q", initialContent, readBuffer)
	}

	// Try to read again, should get EOF or 0 bytes
	secondRead, err := buffer.Read(readBuffer)
	if err != io.EOF {
		t.Errorf("Expected EOF on second read, got %v", err)
	}
	if secondRead != 0 {
		t.Errorf("Expected to read 0 bytes on second read, got %d", secondRead)
	}
}

// TestOurByteBufferReadPartial tests reading content from a buffer in parts
func TestOurByteBufferrReadPartial(t *testing.T) {
	// create a buffer with a test content
	initialContent := []byte("hello world")
	buffer := NewOurByteBuffer(initialContent)

	// create a small destination slice to read just part of the content
	firstReadBuffer := make([]byte, 5)
	// read the first 5 bytes
	n, err := buffer.Read(firstReadBuffer)

	// check for unexpected errors
	if err != nil {
		t.Fatalf("Unexpected error reading from buffer: %v", err)
	}

	// verifty we read the expected numbere of bytes
	if n != 5 {
		t.Errorf("Expected to read 5 bytes, read %d", n)
	}

	// check if what we read matches the first part of the original content
	if !bytes.Equal(firstReadBuffer, []byte("hello")) {
		t.Errorf("Expected to read %q, got %q", "hello", firstReadBuffer)
	}

	//create another destination slice to read the rest of the content
	secondReadBuffer := make([]byte, 10) // larger than needed
	//read the remaining contenet
	n, err = buffer.Read(secondReadBuffer)
	// Check for unexpected errors
	if err != nil && err != io.EOF {
		t.Fatalf("Unexpected error reading from buffer: %v", err)
	}

	// Verify we read the expected number of bytes
	if n != 6 {
		t.Errorf("Expected to read 6 bytes, read %d", n)
	}

	// Check if what we read matches the second part of the original content
	if !bytes.Equal(secondReadBuffer[:n], []byte(" world")) {
		t.Errorf("Expected to read %q, got %q", " world", secondReadBuffer[:n])
	}
}

// TestOurByteBufferBytesAfterRead ensures Bytes() returns all data regardless of read position
func TestOurByteBufferBytesAfterRead(t *testing.T) {
	// create a buffer with test content
	initialContent := []byte("hello world")
	buffer := NewOurByteBuffer(initialContent)

	// Read some bytes to advance the read postion
	readBuffer := make([]byte, 5)
	_, err := buffer.Read(readBuffer)

	// check for unexpected errors
	if err != nil {
		t.Fatalf("Unexpected error reading from buffer: %v", err)
	}

	//check if bytes() still returns all content despite having read some
	result := buffer.Bytes()
	if !bytes.Equal(result, initialContent) {
		t.Errorf("Expected bytes() to return all content %q after reading, got %q",
			initialContent, result)
	}
}
