package js

import (
	"fmt"
	"testing"
)

func Test_JsParser(t *testing.T) {
	filePath := "./sign.js"
	result := JsParser(filePath, "tac")
	fmt.Println(result)
}
