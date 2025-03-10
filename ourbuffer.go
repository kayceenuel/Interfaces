package main

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
