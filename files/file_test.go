package files

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"

	"github.com/edgarSucre/rctlr/utils"
)

func TestOpenInvalidFile(t *testing.T) {
	utils.Info("TestOpenInvalidFile: Tratar de abrir un archivo que no existe returna un error")
	_, err := os.Open("content.txt")
	utils.AssertContains(err.Error(), "content.txt", "Mensaje de error incorrecto", t)
}

func TestOpenFile(t *testing.T) {
	utils.Info(`TestOpenFile: La manera mas facil de abrir un archivo es con la funcion "Open" 
	del packete "os"`)
	f1, err := os.Open("safe.txt")
	utils.AssertTrue(err == nil, `No se pudo abrir el archivo "safe"`, t)

	f2, err := os.Open("./sub/sub.txt")
	utils.AssertTrue(err == nil, `No se pudo abrir el archivo "sub/sub.txt"`, t)

	f3, err := os.Open("../upper.txt")
	utils.AssertTrue(err == nil, `No se pudo abrir el archivo "../upper.txt"`, t)

	f1.Close()
	f2.Close()
	f3.Close()
}

func TestReadFileReadAll(t *testing.T) {
	utils.Info(`TestReadFileReadAll: Una manera de leer el contendio de un archivo es meidantes la
	utilizacion de la funcion ReadALl del paquete ioutil. Esta funcion retorna
	un slice de bytes, el cual debe puede ser transformado a string: "string(content)"`)
	f1, err := os.Open("safe.txt")
	if err != nil {
		utils.Err("No se pudo abrir el archivo safe.txt")
		t.FailNow()
	}
	content, err := ioutil.ReadAll(f1)
	if err != nil {
		utils.Err("No se pudo leer el contenido del archivo safe.txt")
		t.FailNow()
	}
	utils.AssertContains(string(content), "second", "No se encontro el contenido", t)
}

func TestScanLines(t *testing.T) {
	utils.Info(`TestScanFiles: La lectura de archivos de texto, generalmente se hacen 
	tokenizando el texto. Para tokenizar y leer el contenido un token a la vez se utiliza 
	un scanner del paquete bufio`)

	scanner := getScanner(t)
	utils.Info("El tokenizador por default es por linea")
	out := getTokens(scanner)
	utils.AssertEquals("content on safe file", out[0],
		"La primera linea no coincide", t)

	utils.AssertEquals("second line", out[1],
		"La segunda linea no coincide", t)

}

func TestScanWords(t *testing.T) {
	utils.Info(`Para cambiar el scanner por defecto se utiliza la funcion split 
	del paquete bufio`)

	scanner := getScanner(t)
	scanner.Split(bufio.ScanWords)
	out := getTokens(scanner)
	utils.AssertEquals("content", out[0], "La primera palabra no coincide", t)
	utils.AssertEquals("on", out[1], "La segunda palabra no coincide", t)
	utils.AssertEquals("safe", out[2], "La tercera palabra", t)
}

func getTokens(s *bufio.Scanner) []string {
	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

func getScanner(t *testing.T) *bufio.Scanner {
	f, err := os.Open("safe.txt")
	if err != nil {
		utils.Err("No se pudo abrir el archivo safe.txt")
		t.FailNow()
	}
	scanner := bufio.NewScanner(f)
	return scanner
}
