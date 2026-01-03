//go:build js && wasm

package internal

import "syscall/js"

func Doc() js.Value {
	return js.Global().Get("document")
}

func ConsoleLog(v ...any)   { js.Global().Get("console").Call("log", v...) }
func ConsoleWarn(v ...any)  { js.Global().Get("console").Call("warn", v...) }
func ConsoleError(v ...any) { js.Global().Get("console").Call("error", v...) }
