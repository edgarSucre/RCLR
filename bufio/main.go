package main

import "RCTLR/bufio/console"

func main() {
	console.TestEmptyRead()
	console.TestSimpleRead()
	console.TestPeekBuffered()
	console.TestDiscard()
}
