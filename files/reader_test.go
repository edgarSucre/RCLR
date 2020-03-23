package files

import (
	"bufio"
	"os"
	"testing"

	"github.com/edgarSucre/rctlr/utils"
)

func TestReadBytes(t *testing.T) {
	utils.Info(`TestReadRunes: si se requiere un mayor control sobre la lectura, 
	se puede utilizar un bufio.Reader`)
	f := getFile(t)
	defer f.Close()

	reader := bufio.NewReader(f)
	utils.Info(`La manera mas comun de leer con un reader es leer una cantidad de bytes.`)
	out := make([]byte, 7)
	reader.Read(out)
	utils.AssertContains(string(out), "content",
		"El contenido no concuerda", t)

	reader.Read(out)
	utils.AssertContains(string(out), "on",
		"El contenido no concuerda", t)
}

func getFile(t *testing.T) *os.File {
	f, err := os.Open("safe.txt")
	if err != nil {
		utils.Err("No se pudo abrir el archivo safe.txt")
		t.FailNow()
	}
	return f
}
