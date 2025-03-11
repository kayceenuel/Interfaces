package main

import "io"

// Simple Buffer that holds bytes in memory.
//you can write byte to it or read byte from it. Just like Input stuff and output stuff.
type OurByteBuffer struct {
	data []byte //slice of byte that holds all the data (io data.)
}

// NewOurByteBuffer creates a new buffer with some starting text.
// It’s like making a new bucket and pouring in some initial water.
func NewOurByteBuffer(initial string) *OurByteBuffer {
	return &OurByteBuffer{data: []byte(initial)} // Turn the string into bytes and put it in the bucket.
}

// Write adds new bytes to the buffer.
// It’s like pouring more water into the bucket.
func (b *OurByteBuffer) Write(p []byte) (n int, err error) {
	b.data = append(b.data, p...) // Add the new bytes to the end of the bucket’s contents.
	return len(p), nil            // Tell how much was added (size of p) and say “no problems” (nil error).
}

// Read take bytes out of the buffer and puts them into your container.
func (b *OurByteBuffer) Read(p []byte) (n int, err error) {
	if len(b.data) == 0 { // If the bucket is empty...
		return 0, io.EOF // Say “nothing scooped” (0) and “bucket’s dry” (EOF error).
	}
	n = copy(p, b.data) // Scoop as much as fits into your cup (p) from the bucket.
	b.data = b.data[n:] // Remove what you scooped from the bucket.
	return n, nil       // Tell how much you scooped and say “no problems”.
}

// Bytes shows you everything currently in the buffer.
// It’s like looking into the bucket to see what’s there.
func (b *OurByteBuffer) Bytes() []byte {
	return b.data //Just hand over the bucket’s contents.
}
