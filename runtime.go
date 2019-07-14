package quickjs

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lquickjs
// #include <quickjs-libc.h>
import "C"

type JSRuntime struct {
	ref *C.JSRuntime
}

func NewJSRuntime() *JSRuntime {
	rt := new(JSRuntime)
	rt.ref = C.JS_NewRuntime()
	return rt
}