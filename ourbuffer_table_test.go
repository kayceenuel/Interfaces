package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/passes/tests"
)

// TestOurByteBufferTable uses a table-driven approach to test multiple scenarios
func TestOurByteBufferTable(t *testing.T) {
	// Define test case with inputs and expected outputs 
	tests := []struct {
		name 			string 			// Name of the test case 
		initialContent 	[]byte 			// initial buffer content 
		writeContent 	[]byte 			// Content to write (if any) 
		readsize 		int				// size of the read buffer 
		expectedBytes   []byte 			// Expected result from bytes()
		expectedRead 	[]byte 			// Expected read content 
		expectedErr 	error 			// Expected error from Read
	}{
		{
			name:           "Empty buffer",
			initialContent: []byte{},
			readSize:       5,
			expectedBytes:  []byte{},
			expectedRead:   []byte{},
			expectedErr:    io.EOF,

		},
		{
			name: 			"Read less than buffer size",
			initialContent: []byte("hello world"),
			readsize:		3 ,	
			expectedBytes: 	[]byte("hello"),
			expectedRead:	[]byte("hel"),
			expectedErr:	nil,
		},
		{
			name: 			"Read more than buffer size", 
			initialContent: []byte("hello"),
			readsize: 10,
			expectedBytes: 	[]byte("hello"),
			expectedRead: 	[]byte("hello"),
			expectedErr: 	nil, 
		}, 
		{
			name: 			"Write to buffer",
			initialContent: []byte("hello"),
			writeContent: 	[]byte(" world"),
			readsize: 		11,
			expectedBytes: 	[]byte("hello world"),
			expectedRead: 	[]byte("hello world"),
			expectedErr: 	nil,
		}
	}
}