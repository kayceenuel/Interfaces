package main

import (
	"bytes"
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
