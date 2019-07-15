package quickjs

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lquickjs
#include "quickjs-bridge.h"
*/
import "C"

type JSContextRaw struct {
	ref *C.JSContext
}

func JSNewContextRaw(runtime JSRuntime) *JSContextRaw {
	ctx := new(JSContextRaw)
	ctx.ref = C.JS_NewContextRaw(runtime.ref)
	return ctx
}

func (ctx *JSContext) ToContext() JSContext {
	return JSContext{
		ref: ctx.ref,
	}
}

func (ctx *JSContextRaw) AddIntrinsicBaseObjects() {
	C.JS_AddIntrinsicBaseObjects(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicDate() {
	C.JS_AddIntrinsicDate(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicEval() {
	C.JS_AddIntrinsicEval(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicStringNormalize() {
	C.JS_AddIntrinsicStringNormalize(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicRegExpCompiler() {
	C.JS_AddIntrinsicRegExpCompiler(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicRegExp() {
	C.JS_AddIntrinsicRegExp(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicJSON() {
	C.JS_AddIntrinsicJSON(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicProxy() {
	C.JS_AddIntrinsicProxy(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicMapSet() {
	C.JS_AddIntrinsicMapSet(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicTypedArrays() {
	C.JS_AddIntrinsicTypedArrays(ctx.ref)
}

func (ctx *JSContextRaw) AddIntrinsicPromise() {
	C.JS_AddIntrinsicPromise(ctx.ref)
}
