package toutiao

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func Signature(f string) string {
	jsFile := "./signature.js"
	bytes, _ := ioutil.ReadFile(jsFile)
	vm := otto.New()
	vm.Run(string(bytes))
	s, err := vm.Call("get_sign", nil, f)
	if err != nil {
		return " "
	}
	return s.String()
}
