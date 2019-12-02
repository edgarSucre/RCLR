package console

import (
	"bufio"
	"os"

	"github.com/gookit/color"
)

var r *bufio.Reader = bufio.NewReader(os.Stdin)

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

// TestSimpleRead read from the console in to a tiny slice
func TestSimpleRead() {
	color.Info.Println("Reading into an small slice")
	out := make([]byte, 1)
	color.Yellow.Println("Insert some text")
	readed := readToSlice(out)
	if readed != len(out) {
		color.Error.Printf("Expected readed to be %d, but got %d \n", len(out), readed)
	} else {
		color.Cyan.Println("Reading into an small slice worked as expected")
	}
}

// TestPeekBuffered checks whats in the buffer
func TestPeekBuffered() {
	color.Info.Println("Peeking into the buffer")
	out, err := r.Peek(r.Buffered())
	if err != nil {
		color.Error.Println("Could not peek into a console buffer", err)
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
	}
}

// Mega do stuff
func Mega() {
	reader := bufio.NewReader(os.Stdin)
	out := make([]byte, 2)
	reader.Read(out)
	reader.Read(out)
	//reader.Read(out)
	ln, _ := reader.Peek(reader.Buffered())
	color.Warn.Println(string(ln))
	//color.Warn.Println(reader.Buffered())
}

func readToSlice(s []byte) int {
	readed, err := r.Read(s)
	if err != nil {
		color.Error.Println("Failed to read from console", err)
	}
	return readed
}
