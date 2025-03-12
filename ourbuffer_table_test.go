package main

import (
	"bytes"
	"io"
	"testing"
)

// TestOurByteBufferTable uses a table-driven approach to test multiple scenarios
func TestOurByteBufferTable(t *testing.T) {
	// Define test cases with inputs and expected outputs
	tests := []struct {
		name           string // Name of the test case
		initialContent []byte // Initial buffer content
		writeContent   []byte // Content to write (if any)
		readSize       int    // Size of read buffer
		expectedBytes  []byte // Expected result from Bytes()
		expectedRead   []byte // Expected read content
		expectedErr    error  // Expected error from Read
	}{
		{
			name:           "Empty buffer",
			initialContent: []byte{},
			readSize:       5,
			expectedBytes:  []byte{},
			expectedRead:   []byte{},
			expectedErr:    io.EOF,
		},
		{
			name:           "Read exactly buffer size",
			initialContent: []byte("hello"),
			readSize:       5,
			expectedBytes:  []byte("hello"),
			expectedRead:   []byte("hello"),
			expectedErr:    nil,
		},
		{
			name:           "Read less than buffer size",
			initialContent: []byte("hello"),
			readSize:       3,
			expectedBytes:  []byte("hello"),
			expectedRead:   []byte("hel"),
			expectedErr:    nil,
		},
		{
			name:           "Read more than buffer size",
			initialContent: []byte("hello"),
			readSize:       10,
			expectedBytes:  []byte("hello"),
			expectedRead:   []byte("hello"),
			expectedErr:    nil,
		},
		{
			name:           "Write to buffer",
			initialContent: []byte("hello"),
			writeContent:   []byte(" world"),
			readSize:       11,
			expectedBytes:  []byte("hello world"),
			expectedRead:   []byte("hello world"),
			expectedErr:    nil,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer with the test's initial content
			buffer := NewOurByteBuffer(tt.initialContent)

			// If test specifies content to write, do that
			if tt.writeContent != nil {
				_, err := buffer.Write(tt.writeContent)
				if err != nil {
					t.Fatalf("Unexpected error writing to buffer: %v", err)
				}
			}

			// Check if Bytes() returns the expected content
			if !bytes.Equal(buffer.Bytes(), tt.expectedBytes) {
				t.Errorf("Bytes() = %q, want %q", buffer.Bytes(), tt.expectedBytes)
			}

			// Test reading from the buffer
			readBuf := make([]byte, tt.readSize)
			n, err := buffer.Read(readBuf)

			// Check if we got the expected error
			if err != tt.expectedErr && !(err == nil && tt.expectedErr == nil) {
				t.Errorf("Read() error = %v, want %v", err, tt.expectedErr)
			}

			// Check if we read the expected content
			actualRead := readBuf[:n]
			if !bytes.Equal(actualRead, tt.expectedRead) {
				t.Errorf("Read() = %q, want %q", actualRead, tt.expectedRead)
			}
		})
	}
}
