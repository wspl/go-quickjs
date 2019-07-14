package quickjs

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lquickjs
#include "quickjs-bridge.h"
*/
import "C"

//import "unsafe"

//type JSModuleInitFunc func(ctx *JSContext, mod *JSGoModule)

//type JSGoModule struct {
//	ref C.JSModuleDef
//}
//
//func NewJSGoModule(ctx JSContext, moduleName string) {
//	jsm := new(JSGoModule)
//
//	moduleNameCstr := C.CString(moduleName)
//	defer C.free(unsafe.Pointer(moduleNameCstr))
//	jsm.ref = C.JS_NewCModule(ctx.ref, moduleNameCstr)
//}
