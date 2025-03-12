package main

import (
	"io"
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
