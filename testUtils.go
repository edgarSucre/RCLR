package rctlr

import (
	"fmt"
	"testing"
)

const (
	infoColor    = "\033[1;34m\t%s\033[0m\n\n"
	warningColor = "\033[1;33m\t%s\033[0m\n"
	errorColor   = "\033[1;31m\t%s\033[0m\n"
)

func failTest(message string, expected, actual interface{}, t *testing.T) {
	err(message)
	warn(fmt.Sprint("Expected:", expected))
	warn(fmt.Sprint("Actual:", actual))
}

func assertEquals(expected, actual interface{}, message string, t *testing.T) {
	if expected != actual {
		err(message)
		warn(fmt.Sprint("Expected:", expected))
		warn(fmt.Sprint("Actual:", actual, "\n"))
	}
}

func err(m string) {
	fmt.Printf(errorColor, m)
}

func warn(m string) {
	fmt.Printf(warningColor, m)
}

func info(m string) {
	fmt.Printf(infoColor, m)
}
