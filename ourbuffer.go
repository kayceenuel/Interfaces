package main

import "io"

// OurByteBuffer is a custom implementation of a buffer that stores bytes.
// It implements both io.Reader and io.Writer interfaces.
//
// Limitations:
// - Thread Safety: Not safe for concurrent use without external synchronization
// - Memory: Keeps all data in memory until explicitly cleared
// - Size: Limited only by available system memory
type OurByteBuffer struct {
	data    []byte // Stores all the buffer content
	readPos int    // Tracks how much has been read
}

// NewOurByteBuffer creates a new buffer with the given initial data
func NewOurByteBuffer(initialData []byte) *OurByteBuffer {
	// Create and return a new buffer containing the initial data
	return &OurByteBuffer{
		data:    append([]byte{}, initialData...), // Make a copy of the input data
		readPos: 0,                                // Start with nothing read
	}
}

// Write adds more bytes to the end of the buffer
// Implements the io.Writer interface
func (b *OurByteBuffer) Write(p []byte) (n int, err error) {
	// Add the new bytes to the end of our data
	b.data = append(b.data, p...)
	// Return the number of bytes we added (all of them) and no error
	return len(p), nil
}

// Read gets bytes from the buffer into the provided slice
// Implements the io.Reader interface
func (b *OurByteBuffer) Read(p []byte) (n int, err error) {
	// If we've already read everything, return EOF
	if b.readPos >= len(b.data) {
		return 0, io.EOF
	}

	// Calculate how many bytes we can read
	// (either what's left or the size of the destination slice, whichever is smaller)
	bytesRemaining := len(b.data) - b.readPos
	bytesToRead := len(p)
	if bytesToRead > bytesRemaining {
		bytesToRead = bytesRemaining
	}

	// Copy the bytes from our buffer to the destination slice
	n = copy(p, b.data[b.readPos:b.readPos+bytesToRead])

	// Update our reading position
	b.readPos += n

	// Return the number of bytes read and no error
	return n, nil
}

// Bytes returns a slice containing all the bytes in the buffer
func (b *OurByteBuffer) Bytes() []byte {
	// Return a copy of all the data, regardless of read position
	return append([]byte{}, b.data...)
}
