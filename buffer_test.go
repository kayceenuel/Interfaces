package main

import (
	"bytes"
	"testing"
)

func TestBufferWrite(t *testing.T) {
	var b bytes.Buffer
	n, err := b.Write([]byte("hello"))
	if err != nil || n != 5 || b.String() != "hello" {
		t.Errorf("Write failed: n=%d, err=%v, got=%s", n, err, b.String())
	}
}

func TestBufferRead(t *testing.T) {
	b := bytes.NewBufferString("hello")
	p := make([]byte, 5)
	n, err := b.Read(p)
	if err != nil || n != 5 || string(p) != "hello" {
		t.Errorf("Read failed: n=%d, err=%v, got=%s", n, err, p)
	}
}
