package rctlr

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/gookit/color"
)

func TestRuneLiteral(t *testing.T) {
	color.Info.Println(`
	TestRuneLiteral: Go saves individual characters as Unicode.
	This Unicode is represented by the type 'Rune' which is a alias
	for type int32.
	`)
	character := 'E'
	tipo := fmt.Sprint((reflect.TypeOf(character)))
	if tipo != "int32" {
		color.Error.Println("TestRuneLiteral: Characters was not saved as Rune\n")
		logDifferential("int32", tipo)
		t.Fail()
	}
}

func TestStringIteration(t *testing.T) {
	color.Info.Printf("\tTestStringIteration: iterating a string gets you Runes\n")
	phrase := "A collection of words"
	for _, character := range phrase {
		tipo := fmt.Sprint((reflect.TypeOf(character)))
		if tipo != "int32" {
			color.Error.Println("TestStringIteration: Characters was not saved as Rune\n")
			logDifferential("int32", tipo)
			t.Fail()
		}
	}
}

func TestNumberToString(t *testing.T) {
	color.Info.Printf("\tNumberToString: just cast it\n")
	age := 25
	sage := string(age + 30)
	tipo := fmt.Sprint((reflect.TypeOf(sage)))
	if tipo != "string" {
		color.Error.Printf("\tCould not cast number to string\n")
		logDifferential("string", tipo)
		t.Fail()
	}
}

func TestStringToInt(t *testing.T) {
	color.Info.Println(`
	TestStringToInt: From string to int requires the strconv package
	`)
	age := "25"
	sage, err := strconv.ParseInt(age, 10, 32)
	sage += 30
	if sage != 55 || err != nil {
		color.Error.Printf("\tCould not convert string to int\n")
		logDifferential(55, sage)
		t.Fail()
	}
}

func TestStringToFloat(t *testing.T) {
	color.Info.Println(`
	TestStringToInt: From string to float requires the strconv package
	`)
	slim := "80.5"
	fat, err := strconv.ParseFloat(slim, 64)
	fat += 80.0
	if fat != 160.5 || err != nil {
		color.Error.Printf("\tCould not convert string to float\n")
		logDifferential(160.5, fat)
		t.Fail()
	}
}
