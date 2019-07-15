package main

import (
	. "go-quickjs"
)

func main() {
	runtime := NewJSRuntime()
	context := runtime.NewContext()

	fn := context.NewGoFunction(func(args []*JSValue, this *JSValue) (*JSValue, *JSError) {
		println("Invoked!")
		return context.NewString("Hello World"), nil
	})
	fn.Value().Expose("hello")
	ret, err := context.Eval("hello", "")
	if err != nil {
		println(err.Message())
	}
	println(ret.String())
}
