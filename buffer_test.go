package main

import (
	"bytes"
	"testing"
)

func TestBufferWrite(t *testing.T) {
	var b bytes.Buffer
	n, err := b.Write([]byte("Hello"))
	if err != nil || n != 5 || b.String() != "hello" {
		t.Errorf("Write failed: n=%d, err=%v, got=%s", n, err, b.String())
	}
}
