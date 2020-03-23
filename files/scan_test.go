package files

import (
	"bufio"
	"os"
	"testing"

	"github.com/edgarSucre/rctlr/utils"
)

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
	//TODO: close file
}
