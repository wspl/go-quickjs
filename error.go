package quickjs

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lquickjs
#include "quickjs-bridge.h"
*/
import "C"

type JSError struct {
	ref C.JSValue
	ctx *JSContext
	val *JSValue
}

func WrapJSError(ctx *JSContext, cerr C.JSValue) *JSError {
	jse := new(JSError)
	jse.ref = cerr
	jse.ctx = ctx
	return jse
}

func (jse *JSError) Value() *JSValue {
	if jse.val == nil {
		jse.val = jse.ctx.WrapValue(jse.ref)
	}
	return jse.val
}

func (jse *JSError) Message() string {
	return jse.Value().Property("message").String()
}

func (jse *JSError) Errno() string {
	return jse.Value().Property("errno").String()
}
