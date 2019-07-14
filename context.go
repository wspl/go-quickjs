package quickjs

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lquickjs
#include "quickjs-bridge.h"
*/
import "C"
import "unsafe"

var ctxMap = make(map[*C.JSContext]*JSContext)

type JSContext struct {
	ref       *C.JSContext
	functions []*JSGoFunction
	cFunction C.JSValue
	global    *JSValue
}

func NewJSContext(runtime *JSRuntime) *JSContext {
	ctx := new(JSContext)
	ctx.ref = C.JS_NewContext(runtime.ref)
	ctx.functions = []*JSGoFunction{}
	ctx.cFunction = C.JS_NewCFunction(ctx.ref, (*C.JSCFunction)(unsafe.Pointer(C.InvokeGoHandler)), nil, C.int(5))
	ctx.global = ctx.WrapValue(C.JS_GetGlobalObject(ctx.ref))

	ctxMap[ctx.ref] = ctx

	return ctx
}

func (ctx *JSContext) Eval(script string, filename string) (*JSValue, *JSError) {
	scriptCstr := C.CString(script)
	defer C.free(unsafe.Pointer(scriptCstr))
	scriptClen := C.ulong(len(script))

	filenameCstr := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameCstr))

	ret := ctx.WrapValue(C.JS_Eval(ctx.ref, scriptCstr, scriptClen, filenameCstr, C.int(0)))
	e := ctx.Exception()
	if e != nil {
		return ret, e
	}
	return ret, nil
}

func (ctx *JSContext) Global() *JSValue {
	return ctx.global
}

func (ctx *JSContext) NewGoFunction(fn JSGoFunctionCallback) *JSGoFunction {
	return NewJSGoFunction(ctx, fn)
}

func (ctx *JSContext) NewBool(bool bool) *JSValue {
	n := 0
	if bool {
		n = 1
	}
	return ctx.WrapValue(C.JS_NewBool(ctx.ref, C.int(n)))
}

func (ctx *JSContext) NewInt32(int int32) *JSValue {
	return ctx.WrapValue(C.JS_NewInt32(ctx.ref, C.int32_t(int)))
}

func (ctx *JSContext) NewInt64(int int64) *JSValue {
	return ctx.WrapValue(C.JS_NewInt64(ctx.ref, C.int64_t(int)))
}

func (ctx *JSContext) NewFloat64(double float64) *JSValue {
	return ctx.WrapValue(C.JS_NewFloat64(ctx.ref, C.double(double)))
}

func (ctx *JSContext) NewString(string string) *JSValue {
	cstr := C.CString(string)
	defer C.free(unsafe.Pointer(cstr))
	return ctx.WrapValue(C.JS_NewString(ctx.ref, cstr))
}

func (ctx *JSContext) NewNull() *JSValue {
	return ctx.WrapValue(C.JS_NULL)
}

func (ctx *JSContext) NewUndefined() *JSValue {
	return ctx.WrapValue(C.JS_UNDEFINED)
}

func (ctx *JSContext) NewException() *JSValue {
	return ctx.WrapValue(C.JS_EXCEPTION)
}

func (ctx *JSContext) NewUninitialized() *JSValue {
	return ctx.WrapValue(C.JS_UNINITIALIZED)
}

func (ctx *JSContext) WrapValue(cval C.JSValue) *JSValue {
	return NewJSValue(ctx, cval)
}

func (ctx *JSContext) WrapError(cerr C.JSValue) *JSError {
	return WrapJSError(ctx, cerr)
}

func (ctx *JSContext) Exception() *JSError {
	val := ctx.WrapValue(C.JS_GetException(ctx.ref))
	if val.IsNull() {
		return nil
	}
	return val.Error()
}

func (ctx *JSContext) Try() {
	//fn := ctx.NewGoFunction(func(args []*JSValue, this *JSValue) (*JSValue, *JSValue) {
	//	return ctx.NewString("Hello"), nil
	//})
	//ret := fn.Call([]*JSValue {ctx.NewString("Arg1"), ctx.NewBool(true)}, nil)
	//println(ret.String())
	ret, err := ctx.Eval("as", "")
	if err != nil {
		println("1", err.Message())
	}
	ret, err = ctx.Eval("const a = 1", "")
	if err != nil {
		println("2", err.Message())
	}
	println(ret.String())
}
