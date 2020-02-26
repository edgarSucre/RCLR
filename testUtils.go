package rctlr

import "github.com/gookit/color"

func logDifferential(expected, actual interface{}) {
	color.Warn.Println("Expected:", expected)
	color.Warn.Println("Actual:", actual)
}
