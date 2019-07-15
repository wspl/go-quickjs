package main

import "C"
import (
	. "go-quickjs"
)

func main() {
	runtime := NewJSRuntime()
	defer runtime.Free()
	context := runtime.NewContext()
	defer context.Free()

	fn := context.NewGoFunction(func(args []*JSValue, this *JSValue) (*JSValue, *JSError) {
		println("Invoked!")
		return context.NewString("Hello World"), nil
	})
	fn.Value().Expose("hello")
	ret, err := context.Eval("hello()", "")
	if err != nil {
		println(err.Message())
	}
	println(ret.String())
}
