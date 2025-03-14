# Go Interfaces Project
Implementation of custom byte buffer and filtering pipe demonstrating Go interfaces.
## Components

OurByteBuffer: Custom implementation of a bytes buffer (io.Reader and io.Writer)
FilteringPipe: Writer that removes digits from text

## Usage
```bash
// Buffer example
buffer := NewOurByteBuffer([]byte("Hello"))
buffer.Write([]byte(" World"))
fmt.Printf("Content: %s\n", buffer.Bytes())
```
```bash
// Filter example
output := NewOurByteBuffer([]byte(""))
filter := NewFilteringPipe(output)
filter.Write([]byte("abc123def456"))
fmt.Printf("Filtered: %s\n", output.Bytes()) // "abcdef"
```
## RUN

```bash
go run main.go ourbuffer.go filtering_pipe.go
```
## TEST 
```bash
go test -v
```
