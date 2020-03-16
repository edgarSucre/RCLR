package rctlr

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestRuneLiteral(t *testing.T) {
	info(`TestRuneLiteral: Go almacena caracteres individuales como Unicode.
	Unicode es representado por el tipo 'Rune' el cual es un alias para el tipo int32.`)
	character := 'E'
	tipo := fmt.Sprint((reflect.TypeOf(character)))
	assertEquals(tipo, "int32",
		"TestRuneLiteral: Characters was not saved as Rune", t)
}

func TestStringIteration(t *testing.T) {
	info("TestStringIteration: iterando una cadena optienes Runes")
	phrase := "A collection of words"
	for _, character := range phrase {
		tipo := fmt.Sprint((reflect.TypeOf(character)))
		assertEquals(tipo, "int32",
			"TestStringIteration: Caracteres no fueron almacenados como Rune", t)
	}
}

func TestNumberToString(t *testing.T) {
	info("NumberToString: solo es necesario castiarlo")
	age := 25
	sage := string(age + 30)
	tipo := fmt.Sprint((reflect.TypeOf(sage)))
	assertEquals(tipo, "string",
		"NumberToString: no se pudo castear de numero a string", t)
}

func TestStringToInt(t *testing.T) {
	info(`TestStringToInt: De string hacia entero se requiere el usar 
	el paquete strconv y la funcion ParseInt`)
	age := "25"
	sage, _ := strconv.ParseInt(age, 10, 32)
	sage += 30
	assertEquals(sage, int64(55), "Could not convert string to int", t)
}

func TestStringToFloat(t *testing.T) {
	info(`
	TestStringToFloat: De string hacia float se requiere el usar el paquete strconv y
	la funcion ParseFloat`)
	slim := "80.5"
	fat, _ := strconv.ParseFloat(slim, 64)
	fat += 80.0
	assertEquals(fat, float64(160.5), "Could not convert string to float", t)
}

func TestMakeSlice(t *testing.T) {

	info(`TestMake: Go puede crear variables de tipos slice reservar su tamaño usando
	la funcion make`)
	sli := make([]int, 10)
	assertEquals(len(sli), 10,
		"TestMake: no se pudo crear sli con el tamaño especificado", t)

	info("Para slice se puede pasar un tercer parametro para indicar la capacidad > len")
	sli = make([]int, 10, 15)
	assertEquals(cap(sli), 15,
		"TestMake: no se pudo crear sli con la capacidad especificada", t)
}

func TestMakeMap(t *testing.T) {

	info("El valor empty de un map es nil")
	var nilMap map[string]string
	if nilMap != nil {
		failTest("TestMakeMap: el valor de empty de Map no es nil", "nil", nilMap, t)
	}

	makeMap := make(map[string]string)
	info("Map creado con make inicializan los campos con valor zero deacuerdo al tipo")
	assertEquals(makeMap[""], "",
		"TestMake: no se pudo crear makeMap cone el valor zero(string)", t)
}

func TestMakeChan(t *testing.T) {

	makeChan := make(chan int)
	chanType := fmt.Sprint(reflect.TypeOf(makeChan))
	info("Make es la unica manera de crear canales != nil")
	assertEquals(chanType, "chan int",
		"TestMake: no se pudo crear makeChan cone el valor zero(chan it)", t)
}

func TestNew(t *testing.T) {
	newString := new(string)
	newInt := new(int)

	strType := fmt.Sprint(reflect.TypeOf(newString))
	intype := fmt.Sprint(reflect.TypeOf(newInt))

	info("New retorna un pointer del tipo especificado")
	assertEquals(strType, "*string", "TestNew: no se pudo crear un pointer string", t)
	assertEquals(intype, "*int", "TestNew: no se pudo crear un pointer int", t)
}
