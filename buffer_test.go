package main

import (
	"bytes"
	"io"
	"testing"
)

// TestBufferIniialBytes verifies that Bytes() returns initial bytes of a newly creaated buffer
func TestBufferInitialBytes(t *testing.T) {
	// create a buffer with content
	buf := bytes.NewBufferString("hello")
	// checks if Bytes matches content ("hello")
	expected := []byte("hello")
	if got := buf.Bytes(); !bytes.Equal(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

// TestBufferWrite test writing to the Buffer
func TestBufferWrite(t *testing.T) {
	buff := bytes.NewBufferString("hello")
	_, err := buff.Write([]byte(" world"))
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	expected := []byte("hello world")
	if got := buff.Bytes(); !bytes.Equal(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}

// TestBufferReadAll ensures that reading all bytes at once retrieves the full content and empties the buffer.
func TestBufferReadAll(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	p := make([]byte, 11)
	n, err := buf.Read(p)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if n != 11 {
		t.Errorf("Expected to read 11 bytes, got %d", n)
	}
	if !bytes.Equal(p, []byte("hello world")) {
		t.Errorf("Expected %v, got %v", []byte("hello world"), p)
	}
	// Verify the buffer is empty after reading all bytes
	n, err = buf.Read(p)
	if n != 0 || err != io.EOF {
		t.Errorf("Expected 0, io.EOF; got %d, %v", n, err)
	}
}

// TestBufferReadPartial tests reading bytes partially
func TestBufferReadPartial(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	p1 := make([]byte, 5)
	n, err := buf.Read(p1)
	if err != nil {
		t.Errorf("Expected to read 5 bytes, got %v", err)
	}
	if n != 5 {
		t.Errorf("Expected to read 5 bytes, got %d", n)
	}
	if !bytes.Equal(p1, []byte("hello")) {
		t.Errorf("Expected %v, got %v", []byte("hello"), p1)
	}

}
