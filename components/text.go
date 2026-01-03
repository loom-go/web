//go:build js && wasm

package components

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
	"github.com/AnatoleLucet/loom-web/internal"
	. "github.com/AnatoleLucet/loom/components"
	. "github.com/AnatoleLucet/loom/signals"
)

func Text(content any) Node {
	return NodeFunc(func(ctx *RenderContext) error {
		parent := ctx.Get("parent").(js.Value)

		// update
		if ctx.Get("self") != nil {
			self := ctx.Get("self").(js.Value)
			self.Set("nodeValue", content)
			return nil
		}

		// mount
		self := internal.Doc().Call("createTextNode", content)
		ctx.Set("self", self)
		parent.Call("appendChild", self)

		// cleanup
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
