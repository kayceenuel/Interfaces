package main

import (
	"fmt"
)

func main() {
	// Demostrate OurByBuffer
	fmt.Println("=== OurByteBuffer Demo ===")

	// Create a new buffer with initial content
	buffer := NewOurByteBuffer([]byte("Hello"))
	fmt.Printf("Initial buffer: %s\n", buffer.Bytes())

}
