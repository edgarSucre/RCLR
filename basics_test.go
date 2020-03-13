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
	assertEquals(tipo, "int32", "TestRuneLiteral: Characters was not saved as Rune", t)
}

func TestStringIteration(t *testing.T) {
	color.Info.Println("\tTestStringIteration: iterando una cadena optienes Runes\n")
	phrase := "A collection of words"
	for _, character := range phrase {
		tipo := fmt.Sprint((reflect.TypeOf(character)))
		assertEquals(tipo, "int32", "TestStringIteration: Caracteres no fueron almacenados como Rune", t)
	}
}

func TestNumberToString(t *testing.T) {
	color.Info.Println("\tNumberToString: solo es necesario castiarlo")
	age := 25
	sage := string(age + 30)
	tipo := fmt.Sprint((reflect.TypeOf(sage)))
	assertEquals(tipo, "string", "\tNumberToString: no se pudo castear de numero a string", t)
}

func TestStringToInt(t *testing.T) {
	color.Info.Print(`
	TestStringToInt: De string hacia entero se requiere el usar el paquete strconv y la funcion ParseInt
	`)
	age := "25"
	sage, _ := strconv.ParseInt(age, 10, 32)
	sage += 30
	assertEquals(sage, int64(55), "\tCould not convert string to int\n", t)
}

func TestStringToFloat(t *testing.T) {
	color.Info.Println(`
	TestStringToFloat: De string hacia float se requiere el usar el paquete strconv y la funcion ParseFloat
	`)
	slim := "80.5"
	fat, _ := strconv.ParseFloat(slim, 64)
	fat += 80.0
	assertEquals(fat, float64(160.5), "\tCould not convert string to float\n", t)
}

func TestMakeSlice(t *testing.T) {

	color.Info.Println(`
	TestMake: Go puede crear variables de tipos slice reservar su tamaño usando la funcion make
	`)
	sli := make([]int, 10)
	assertEquals(len(sli), 10, "\tTestMake: no se pudo crear sli con el tamaño especificado\n", t)

	color.Info.Println("\tpara slice se puede pasar un tercer parametro para indicar la capacidad > len")
	sli = make([]int, 10, 15)
	assertEquals(cap(sli), 15, "\tTestMake: no se pudo crear sli con la capacidad especificada\n", t)
}

func TestMakeMap(t *testing.T) {

	color.Info.Println("\tEl valor empty de un map es nil")
	var nilMap map[string]string
	if nilMap != nil {
		failTest("\tTestMakeMap: el valor de empty de Map no es nil\n", "nil", nilMap, t)
	}

	makeMap := make(map[string]string)
	color.Info.Println("\tMap creado con make inicializan los campos con valor zero deacuerdo al tipo")
	assertEquals(makeMap[""], "", "\tTestMake: no se pudo crear makeMap cone el valor zero(string)\n", t)
}

func TestMakeChan(t *testing.T) {

	makeChan := make(chan int)
	chanType := fmt.Sprint(reflect.TypeOf(makeChan))
	color.Info.Println("\tmake es la unica manera de crear canales != nil")
	assertEquals(chanType, "chan int", "\tTestMake: no se pudo crear makeChan cone el valor zero(chan it)\n", t)
}

//TODO: crear test usando la funcion new
