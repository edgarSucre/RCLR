package console

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadWithEmptySlice(t *testing.T) {
	var out []byte
	r := getReader("this is a test")
	read(r, out, t)
	if len(out) > 0 {
		t.Errorf("Expected len(out) to be 0, got %d instead", len(out))
	}
}

func TestReadWithNormalSlice(t *testing.T) {
	out := make([]byte, 20)
	input := "this is a test"
	r := getReader(input)
	n := read(r, out, t)
	so := string(out[:n])
	if so != input {
		t.Errorf("Expected so to be '%v', got '%v' instead", input, so)
	}
}

func TestBufferedPeek(t *testing.T) {
	out := make([]byte, 4)
	input := "this is a test"
	r := getReader(input)
	read(r, out, t)
	so := string(out)
	if so != "this" {
		t.Errorf("Expected so to be '%v', got '%v' instead", "this", so)
	}

	unread, err := r.Peek(r.Buffered())
	if err != nil {
		t.Error("failed to peek the buffer")
	}

	tMissing := " is a test"
	if string(unread) != tMissing {
		t.Errorf("Expected unread to be '%v', got '%v' instead", tMissing, string(unread))
	}
}

func read(r *bufio.Reader, out []byte, t *testing.T) int {
	n, err := r.Read(out)
	if err != nil {
		t.Error("Could not read from reader")
	}
	return n
}

func getReader(s string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(s))
}
