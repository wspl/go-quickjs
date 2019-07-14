#include "_cgo_export.h"

JSValue InvokeGoHandler(JSContext *ctx, JSValueConst thisObj, int argc, JSValueConst *argv) {
	return GoHandler(ctx, thisObj, argc, argv);
}