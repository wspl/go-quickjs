package main

import (
	"go-quickjs"
)

func main() {
	runtime := quickjs.NewJSRuntime()
	defer runtime.Free()
	context := runtime.NewContext()
	defer context.Free()

	ret, err := context.Eval("'Hello ' + 'World!'", "")
	if err != nil {
		println(err.Message())
	}
	println(ret.String())
}
