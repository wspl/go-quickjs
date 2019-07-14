package quickjs

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lquickjs
#include "quickjs-bridge.h"
*/
import "C"
import "unsafe"

//export GoHandler
func GoHandler(cctx *C.JSContext, cthis C.JSValueConst, cargc C.int, cargv *C.JSValueConst) C.JSValue {
	ctx := ctxMap[cctx]
	cargs := (*[1 << 28]C.JSValueConst)(unsafe.Pointer(cargv))[:cargc:cargc]
	id := ctx.WrapValue(cargs[0]).Int()
	this := ctx.WrapValue(cthis)
	var args []*JSValue
	for _, carg := range cargs {
		args = append(args, ctx.WrapValue(carg))
	}
	result, err := ctx.functions[id].fn(args, this)
	if err != nil {
		return err.ref
	}
	return result.ref
}

type JSGoFunctionCallback func(args []*JSValue, this *JSValue) (*JSValue, *JSError)

type JSGoFunction struct {
	ref C.JSValue
	ctx *JSContext
	fn  JSGoFunctionCallback
}

func NewJSGoFunction(ctx *JSContext, fn JSGoFunctionCallback) *JSGoFunction {
	jsf := new(JSGoFunction)
	jsf.ctx = ctx
	jsf.fn = fn
	jsf.init()
	return jsf
}

func (jsf *JSGoFunction) init () {
	wrapperScript := "(invokeGoFunction, id) => function () { return invokeGoFunction.call(this, id, arguments) }"
	wrapperFn, _ := jsf.ctx.Eval(wrapperScript, "")

	id := len(jsf.ctx.functions)
	jsf.ctx.functions = append(jsf.ctx.functions, jsf)

	wrapperArgs := []C.JSValue{
		jsf.ctx.cFunction,
		jsf.ctx.NewInt32(int32(id)).ref,
	}
	jsf.ref = C.JS_Call(jsf.ctx.ref, wrapperFn.ref, jsf.ctx.NewNull().ref, 2, &wrapperArgs[0])
}

func (jsf *JSGoFunction) Value () *JSValue {
	return jsf.ctx.WrapValue(jsf.ref)
}

func (jsf *JSGoFunction) Call (args []*JSValue, this *JSValue) *JSValue {
	return jsf.Value().Call(args, this)
}