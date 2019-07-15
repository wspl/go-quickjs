package main

import (
	. "go-quickjs"
)

func main() {
	runtime := NewJSRuntime()
	context := runtime.NewContext()

	context.Try()

	context.Free()
	runtime.Free()
}
