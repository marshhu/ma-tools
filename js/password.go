package js

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func PasswordEnc(salt string, password string) string {
	jsFile := "./password.js"
	bytes, _ := ioutil.ReadFile(jsFile)
	vm := otto.New()
	vm.Run(string(bytes))
	enc, err := vm.Call("password_enc", nil, salt, password)
	if err != nil {
		return " "
	}
	return enc.String()
}
