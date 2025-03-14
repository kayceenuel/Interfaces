package main

import (
	"fmt"
	"os"
)

func main() {
	// Demostrate OurByBuffer
	fmt.Println("=== OurByteBuffer Demo ===")

	// Create a new buffer with initial content
	buffer := NewOurByteBuffer([]byte("Hello"))
	fmt.Printf("Initial buffer: %s\n", buffer.Bytes())

	// Write to the buffer
	buffer.Write([]byte(", world"))
	fmt.Printf("After writing: %s\n", buffer.Bytes())

	//Read from the buffer
	readBuf := make([]byte, 5)
	n, _ := buffer.Read(readBuf)
	fmt.Printf("Read %d bytes: %s\n", n, readBuf[:n])

	// Check what's left inn the buffer
	fmt.Printf("Bytes still in buffer: %s\n", buffer.Bytes())

	// Demostrate FilteringPipe
	fmt.Println("\n=== FilteringPipe Demo ===")

	// create a filtering pipe that writes to standard outputs
	filter := NewFilteringPipe(os.Stdout)

	// Write a string with digits to the filter
	fmt.Print("Filtering result: ")
	filter.Write([]byte("abc2123def456"))
	fmt.Println() // Add a newline
}
