package console

import (
	"bufio"
	"os"

	"github.com/gookit/color"
)

var r *bufio.Reader = bufio.NewReader(os.Stdin)

func init() {
	color.Info.Println("**************************************************")
	color.Info.Println("************** Testing bufio Reader **************")
	color.Info.Println("**************************************************")
}

// TestEmptyRead read nothing since the len(slice) = 0
func TestEmptyRead() {
	color.Info.Println("Reading into an empty slice should do nothing")
	var out []byte
	readed := readToSlice(out)
	if readed != 0 {
		color.Error.Println("Expected readed to be 0, but got %d \n", readed)
	}
	color.Cyan.Println("Reading into an empty slice worked as expected")
}

// TestSimpleRead read from the console in to a slice
func TestSimpleRead() {
	color.Info.Println("With a bigger buffer we can read more text")
	out := make([]byte, 100)
	color.Yellow.Println("Insert some text max(100)")
	readToSlice(out)
	color.Cyan.Println("readed:", string(out))
}

// TestPeekBuffered checks whats in the buffer
func TestPeekBuffered() {
	out := make([]byte, 1)
	color.Info.Println("With a smaller buffer (not 0) the reader put the rest on the buffer")
	color.Yellow.Println("Insert some words")
	readToSlice(out)
	if len(out) > 1 {
		color.Error.Printf("Expected io buffer to be len 1 got %d instead", len(out))
	}

	color.Info.Println("Peeking into the buffer")
	out, err := r.Peek(r.Buffered())
	if err != nil {
		color.Error.Println("Could not peek into a console buffer", err)
	}

	if len(out) != r.Buffered() {
		color.Error.Printf("Expected bufio buffer to be len %d got %d instead", r.Buffered(), len(out))
	}
	color.Cyan.Printf("Peeked from the buffer: %v\n", string(out))
}

// TestDiscard tesit if discard clean the buffer
func TestDiscard() {
	td := r.Buffered() - 1
	color.Info.Printf("Discarting %d bytes from the buffer\n", td)
	n, err := r.Discard(td)
	if err != nil {
		color.Error.Println("Could not discard from the console buffer", err)
	} else if n != td {
		color.Error.Printf("Expected 0 bytes on the buffer, got %d\n", n)
	} else {
		color.Cyan.Println("Discard worked as expected")
	}
}

func readToSlice(s []byte) int {
	readed, err := r.Read(s)
	if err != nil {
		color.Error.Println("Failed to read from console", err)
	}
	return readed
}
