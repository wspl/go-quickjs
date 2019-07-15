package quickjs

/*
#include "quickjs-bridge.h"
*/
import "C"

type JSRuntime struct {
	ref *C.JSRuntime
}

func NewJSRuntime() *JSRuntime {
	rt := new(JSRuntime)
	rt.ref = C.JS_NewRuntime()
	return rt
}

func (rt *JSRuntime) NewContext() *JSContext {
	return NewJSContext(rt)
}

func (rt *JSRuntime) Free() {
	C.JS_FreeRuntime(rt.ref)
}
