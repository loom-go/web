//go:build js && wasm

package web

import "syscall/js"

func doc() js.Value {
	return js.Global().Get("document")
}
