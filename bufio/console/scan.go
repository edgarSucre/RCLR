package console

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/gookit/color"
)

func init() {
	color.Info.Println("**************************************************")
	color.Info.Println("************* Testing bufio Scanner **************")
	color.Info.Println("**************************************************")
}

var (
	singleLineInput string = "this is a test"
	multiLineInput  string = `
		this is also a test
		that have more than one line
	`
)

// TestWordScanner returns all words
func TestWordScanner() {
	r := strings.NewReader(singleLineInput)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		color.Error.Println("The scanner could not finish")
	}
}
