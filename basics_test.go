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
	TestRuneLiteral: Go almacena caracteres individuales como Unicode.
	Unicode es representado por el tipo 'Rune' el cual es un alias para el tipo int32.
	`)
	character := 'E'
	tipo := fmt.Sprint((reflect.TypeOf(character)))
	if tipo != "int32" {
		color.Error.Println("TestRuneLiteral: Characters was not saved as Rune")
		logDifferential("int32", tipo)
		t.Fail()
	}
}

func TestStringIteration(t *testing.T) {
	color.Info.Println("\tTestStringIteration: iterando una cadena optienes Runes\n")
	phrase := "A collection of words"
	for _, character := range phrase {
		tipo := fmt.Sprint((reflect.TypeOf(character)))
		if tipo != "int32" {
			color.Error.Println("TestStringIteration: Caracteres no fueron almacenados como Rune")
			logDifferential("int32", tipo)
			t.Fail()
		}
	}
}

func TestNumberToString(t *testing.T) {
	color.Info.Println("\tNumberToString: solo es necesario castiarlo")
	age := 25
	sage := string(age + 30)
	tipo := fmt.Sprint((reflect.TypeOf(sage)))
	if tipo != "string" {
		color.Error.Println("\tNumberToString: no se pudo castear de numero a string")
		logDifferential("string", tipo)
		t.Fail()
	}
}

func TestStringToInt(t *testing.T) {
	color.Info.Print(`
	TestStringToInt: De string hacia entero se requiere el usar el paquete strconv y la funcion ParseInt
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
	TestStringToFloat: De string hacia float se requiere el usar el paquete strconv y la funcion ParseFloat
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

func TestMake(t *testing.T) {

	color.Info.Println(`
	TestMake: Go puede crear variables de tipos slice, map, chan y reservar su tamaño usando la funcion make
	`)

	makeSlice := make([]int, 10)
	makeMap := make(map[string]string)
	makeChan := make(chan int)

	color.Info.Println("\tmake se puede usar para general slices dinamicos")
	if len(makeSlice) != 10 {
		color.Error.Printf("\tTestMake: no se pudo crear makeSlice con el tamaño especificado\n")
		logDifferential(10, len(makeSlice))
		t.FailNow()
	}

	color.Info.Println("\tpara slice se puede pasar un tercer parametro para indicar la capacidad > len")
	makeSlice = make([]int, 10, 15)
	if cap(makeSlice) != 15 {
		color.Error.Printf("\tTestMake: no se pudo crear makeSlice con la capacidad especificada\n")
		logDifferential(15, cap(makeSlice))
		t.FailNow()
	}

	color.Info.Println("\tmap creado con make inicializan los campos con valor zero deacuerdo al tipo")
	if makeMap[""] != "" {
		color.Error.Printf("\tTestMake: no se pudo crear makeMap cone el valor zero(string)\n")
		logDifferential("Cadena Vacia", makeMap)
		t.FailNow()
	}

	chanType := fmt.Sprint(reflect.TypeOf(makeChan))

	color.Info.Println("\tmake es la unica manera de crear canales != nil")
	if chanType != "chan int" || makeChan == nil {
		color.Error.Printf("\tTestMake: no se pudo crear makeChan cone el valor zero(chan it)\n")
		logDifferential("chan int", chanType)
		t.FailNow()
	}
}
