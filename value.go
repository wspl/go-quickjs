package quickjs

import "unsafe"

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lquickjs
#include "quickjs-bridge.h"
*/
import "C"

var valueMap = make(map[*C.JSValue]JSValue)

type JSValue struct {
	ref C.JSValue
	ctx *JSContext
}

func NewJSValue(ctx *JSContext, cval C.JSValue) *JSValue {
	val := new(JSValue)
	val.ref = cval
	val.ctx = ctx
	return val
}

func (val *JSValue) Free() {
	C.JS_FreeValue(val.ctx.ref, val.ref)
}

func (val *JSValue) AttachContext(ctx *JSContext) {
	val.ctx = ctx
}

func (val *JSValue) String() string {
	cstr := C.JS_ToCString(val.ctx.ref, val.ref)
	return C.GoString(cstr)
}

func (val *JSValue) Float64() float64 {
	i := C.double(0)
	C.JS_ToFloat64(val.ctx.ref, &i, val.ref)
	return float64(i)
}

func (val *JSValue) Int64() int64 {
	i := C.int64_t(0)
	C.JS_ToInt64(val.ctx.ref, &i, val.ref)
	return int64(i)
}

func (val *JSValue) Int() int {
	i := C.int64_t(0)
	C.JS_ToInt64(val.ctx.ref, &i, val.ref)
	return int(i)
}

func (val *JSValue) Int32() int32 {
	i := C.int32_t(0)
	C.JS_ToInt32(val.ctx.ref, &i, val.ref)
	return int32(i)
}

func (val *JSValue) Bool() bool {
	i := C.JS_ToBool(val.ctx.ref, val.ref)
	return val.valueToBool(i)
}

func (val *JSValue) Error() *JSError {
	return val.ctx.WrapError(val.ref)
}

func (val *JSValue) valueToBool(cval C.int) bool {
	if cval == 1 {
		return true
	}
	return false
}

func (val *JSValue) IsInteger() bool {
	return val.valueToBool(C.JS_IsInteger(val.ref))
}

func (val *JSValue) IsBigFloat() bool {
	return val.valueToBool(C.JS_IsBigFloat(val.ref))
}

func (val *JSValue) IsBool() bool {
	return val.valueToBool(C.JS_IsBool(val.ref))
}

func (val *JSValue) IsUndefined() bool {
	return val.valueToBool(C.JS_IsUndefined(val.ref))
}

func (val *JSValue) IsNull() bool {
	return val.valueToBool(C.JS_IsNull(val.ref))
}

func (val *JSValue) IsException() bool {
	return val.valueToBool(C.JS_IsException(val.ref))
}

func (val *JSValue) IsUninitialized() bool {
	return val.valueToBool(C.JS_IsUninitialized(val.ref))
}

func (val *JSValue) IsString() bool {
	return val.valueToBool(C.JS_IsString(val.ref))
}

func (val *JSValue) IsSymbol() bool {
	return val.valueToBool(C.JS_IsSymbol(val.ref))
}

func (val *JSValue) IsObject() bool {
	return val.valueToBool(C.JS_IsObject(val.ref))
}

func (val *JSValue) IsError() bool {
	return val.valueToBool(C.JS_IsError(val.ctx.ref, val.ref))
}

func (val *JSValue) IsFunction() bool {
	return val.valueToBool(C.JS_IsFunction(val.ctx.ref, val.ref))
}

func (val *JSValue) IsConstructor() bool {
	return val.valueToBool(C.JS_IsConstructor(val.ctx.ref, val.ref))
}

func (val *JSValue) IsExtensible() bool {
	return val.valueToBool(C.JS_IsExtensible(val.ctx.ref, val.ref))
}

func (val *JSValue) IsInstanceOf(ctor *JSValue) bool {
	return val.valueToBool(C.JS_IsInstanceOf(val.ctx.ref, val.ref, ctor.ref))
}

func (val *JSValue) Call(args []*JSValue, this *JSValue) *JSValue {
	if this == nil {
		this = val.ctx.Global()
	}
	var cargs []C.JSValue
	for _, arg := range args {
		cargs = append(cargs, arg.ref)
	}
	return val.ctx.WrapValue(C.JS_Call(val.ctx.ref, val.ref, this.ref, C.int(len(args)), &cargs[0]))
}

func (val *JSValue) PropertyByInt(key int) *JSValue {
	return val.ctx.WrapValue(C.JS_GetPropertyUint32(val.ctx.ref, val.ref, C.uint32_t(key)))
}

func (val *JSValue) Property(key string) *JSValue {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	return val.ctx.WrapValue(C.JS_GetPropertyStr(val.ctx.ref, val.ref, cstr))
}

func (val *JSValue) SetPropertyByInt(key int, value *JSValue) {
	C.JS_SetPropertyInt64(val.ctx.ref, val.ref, C.int64_t(key), value.ref)
}

func (val *JSValue) SetProperty(key string, value *JSValue) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	C.JS_SetPropertyStr(val.ctx.ref, val.ref, cstr, value.ref)
}

func (val *JSValue) Expose(name string) {
	val.ctx.global.SetProperty(name, val)
}
