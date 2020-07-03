package js

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func JsParser(filePath string, functionName string, args ...interface{}) (result string) {
	//读入文件
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()
	script, err := vm.Compile("", bytes)
	if err != nil {
		panic(err)
	}
	_, err = vm.Run(script)
	if err != nil {
		panic(err)
	}
	value, err := vm.Call(functionName, nil, args...)
	if err != nil {
		panic(err)
	}

	return value.String()
}
