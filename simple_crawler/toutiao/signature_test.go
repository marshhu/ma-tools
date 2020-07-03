package toutiao

import (
	"fmt"
	"testing"
)

func Test_Signature(t *testing.T) {
	f := "05ef884c100927e0fb59e"
	s := Signature(f)
	fmt.Println(s)
}
