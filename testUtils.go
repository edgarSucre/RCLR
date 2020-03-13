package rctlr

import (
	"testing"

	"github.com/gookit/color"
)

func failTest(message string, expected, actual interface{}, t *testing.T) {
	color.Error.Printf(message)
	color.Warn.Println("Expected:", expected)
	color.Warn.Println("Actual:", actual)
}

func assertEquals(expected, actual interface{}, message string, t *testing.T) {
	if expected != actual {
		color.Error.Printf(message)
		color.Warn.Println("Expected:", expected)
		color.Warn.Println("Actual:", actual)
	}
}
