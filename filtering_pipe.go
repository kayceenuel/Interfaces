package main

import (
	"io"
	"unicode"
)

//filteringPipe filters out digit characters when writing
// to the underlying writer.
//
// Limitations:
// - Thread Saftey: Not safe for concurrent writes without external synchronization
// - Counting: returns the original input length, filtered length.
// - filtering: Only filters digits (0-9), not other numeric characters

type filteringPipe struct {
	writer io.Writer // The underlying writer to send filtered content to
}

// NewFilteringPipe creates a new filtering pipe that writes to the provided writer
func NewFilteringPipe(w io.Writer) *filteringPipe {
	//Create and return a new FilteringPipe that will write to w
	return &filteringPipe{writer: w}
}

// write filters out digits from the input writes the result
// To the underlying writer. Returns the original input length
// Implements the io.writer interface.
func (fp *filteringPipe) Write(p []byte) (n int, err error) {
	//Create a new slice to the filtered content
	filtered := make([]byte, 0, len(p))

	//Go through each byte in the input
	for _, b := range p {
		// / if it's not a digit, add it to our filtered slice
		if !unicode.IsDigit(rune(b)) {
			filtered = append(filtered, b)
		}
	}

	// If there's nothing left after filtering, just retunrn success
	if len(filtered) == 0 {
		return len(p), nil
	}

	// Write the filtered content, to the underlying writer
	_, err = fp.writer.Write(filtered)

	// Return the original length and any error from the underlying write
	return len(p), err
}
