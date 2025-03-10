package main

import "io"

type OurByteBuffer struct {
	data []byte
}

// constructor
func NewOurByteBuffer(initial string) *OurByteBuffer {
	return &OurByteBuffer{data: []byte(initial)}
}

func (b *OurByteBuffer) Write(p []byte) (n int, err error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

func (b *OurByteBuffer) Read(p []byte) (n int, err error) {
	if len(b.data) == 0 {
		return 0, io.EOF
	}
	n = copy(p, b.data)
	b.data = b.data[n:]
	return n, nil
}

func (b *OurByteBuffer) Bytes() []byte {
	return b.data
}
