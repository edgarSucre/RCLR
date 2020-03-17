package files

import (
	"os"
	"testing"

	"github.com/edgarSucre/rctlr/utils"
)

func TestOpenInvalidFile(t *testing.T) {
	utils.Info("Tratar de abrir un archivo que no existe returna un error")
	_, err := os.Open("content.txt")
	utils.AssertContains(err.Error(), "content.txt", "Mensaje de error incorrecto", t)
}

func TestOpenFileSameFolder(t *testing.T) {
	utils.Info(`La manera mas facil de abrir un archivo es con la funcion "Open" 
	del packete "os"`)
	_, err := os.Open("safe.txt")
	utils.AssertTrue(err == nil, `No se pudo abrir el archivo "safe"`, t)
}
