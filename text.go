//go:build js && wasm

package web

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
	. "github.com/AnatoleLucet/loom/components"
	. "github.com/AnatoleLucet/loom/signals"
)

func Text(content any) Node {
	return NodeFunc(func(ctx *RenderContext) error {
		parent := ctx.Get("parent").(js.Value)

		// if the node already exists, just update its content
		if ctx.Get("self") != nil {
			self := ctx.Get("self").(js.Value)
			self.Set("nodeValue", content)
			return nil
		}

		self := doc().Call("createTextNode", content)
		ctx.Set("self", self)
		parent.Call("appendChild", self)

		OnCleanup(func() {
			parent.Call("removeChild", self)
		})

		return nil
	})
}

func BindText[T any](value func() T) Node {
	return Bind(func() Node {
		return Text(value())
	})
}
