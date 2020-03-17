package files

import (
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
	utils.Info(`Una manera de leer el contendio de un archivo es meidantes la
	utilizacion de la funcion ReadALl del paquete ioutil`)
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
