# go-quickjs
QuickJS engine bindings for Go.

**Warning: At present, both the original project `quickjs` and this project are still in the early stage of development. Please use this project carefully in the production environment.**

## Features
* Eval script
* Eval bytecode in `[]byte`
* Compile bytecode into `[]byte`
* Simple exception throwing and catching
* Invoke Go function from JavaScript
* Operate JavaScript values and objects in Go

## Get Started
Install: (MacOS tested only)
```bash
wget https://raw.github.com/wspl/go-quickjs/master/install.sh && sh ./install.sh
```
Hello world:
```go
package main

import "github.com/wspl/go-quickjs"

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
```
Invoke Go function in JavaScript:
```go
package main

import . "github.com/wspl/go-quickjs"

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

```

## TODOs
- Test cases
- Module support
- Fix bugs

## License
[MIT](./LICENSE)