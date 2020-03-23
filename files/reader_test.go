package files

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/edgarSucre/rctlr/utils"
)

func TestReaderBehaviour(t *testing.T) {
	utils.Info(`TestReaderBehaviour: si se requiere un mayor control sobre la 
	lectura, se puede utilizar bufio.Reader. Este reader provee la funcionalidad de
	io.Reader con un buffer y funcionalidades añadidas`)
	f := getFile(t)
	defer f.Close()

	reader := bufio.NewReader(f)
	utils.Info(`La manera mas comun de leer con un reader es usar la función read`)
	out := make([]byte, 7)
	readed, _ := reader.Read(out)

	utils.AssertContains(string(out), "content",
		"El contenido no concuerda", t)

	utils.Info(`Esta función lee y guarda una porcion del texto no leido en el buffer.
	La capacidad del buffer por default es de 4096 bytes así que una ves leido, el
	reader llenara el buffer con el texto que quepa.`)

	utils.Info("La cantidad de bytes leidos debe ser igual al tamaño del slice")
	utils.AssertEquals(readed, 7, "No se leyo la cantidad esperada", t)

	utils.Info("El tamaño del archivo(33) menos los bytes leidos(7) debe ser el buffer")
	utils.AssertEquals(26, reader.Buffered(), "El tamaño del buffer no concuerda", t)

	reader.Read(out)

	utils.AssertContains(string(out), "on",
		"El contenido no concuerda", t)

	utils.Info("Se volvio a leer, asi ql el buffer debe ser menor (26-7)")
	utils.AssertEquals(19, reader.Buffered(), "El tamaño del buffer no concuerda", t)
}

func TestReaderPeek(t *testing.T) {
	utils.Info(`TestReaderPeek: Peek nos permite ver el contenido del buffer sin 
	avanzar el reader`)
	f := getFile(t)
	defer f.Close()

	out := make([]byte, 8)
	reader := bufio.NewReader(f)
	reader.Read(out)

	utils.Info("A la función Peek se le pasa la cantidad de bytes a revisar")
	buffer := reader.Buffered()
	peeked, _ := reader.Peek(2)
	utils.AssertEquals("on", string(peeked), "Lo examinado no concuerda", t)
	utils.AssertEquals(buffer, reader.Buffered(), "Reader fue avanzado", t)
}

func TestReaderDiscard(t *testing.T) {
	utils.Info(`TestReaderPeek: Discard nos permite descartar lo que esta en el buffer`)

	f := getFile(t)
	defer f.Close()

	out := make([]byte, 8)
	reader := bufio.NewReader(f)
	reader.Read(out)

	utils.Info("Primero examinamos el buffer para ver el token existe")
	token, _ := reader.Peek(7)
	utils.AssertEquals("on safe", string(token), "El buffer no contiene el token", t)

	utils.Info("Descartamos la misma cantidad de bytes, y esaminamos el buffer denuevo")
	reader.Discard(7)
	peeked, _ := reader.Peek(25)
	utils.AssertNotContains(string(peeked), "on safe",
		"El buffer aun contiene token descartado", t)
}

func TestReadRuneVSReadByte(t *testing.T) {
	sr := strings.NewReader("日")
	reader := bufio.NewReader(sr)

	utils.Info(`TestReadRuneVSReadByte: read byte returns the first byte
	representing UTF-8 encodend string`)
	by, _ := reader.ReadByte()
	utils.AssertEquals(by, byte(230), "El byte leido no concuerda", t)
	reader.UnreadByte()

	utils.Info(`TestReadRuneVSReadByte: ReadRune returns the rune representation
	of a character of the string`)
	ru, _, _ := reader.ReadRune()
	utils.AssertEquals(ru, rune(26085), "Rune no concuerda", t)
}

func getFile(t *testing.T) *os.File {
	f, err := os.Open("safe.txt")
	if err != nil {
		utils.Err("No se pudo abrir el archivo safe.txt")
		t.FailNow()
	}
	return f
}
