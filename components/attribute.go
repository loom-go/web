//go:build js && wasm

package components

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
	. "github.com/AnatoleLucet/loom/components"
	. "github.com/AnatoleLucet/loom/signals"
)

func Attr(name, value any) Node {
	return NodeFunc(func(ctx *RenderContext) error {
		parent := ctx.Get("parent").(js.Value)

		// short path for properties (might want to create another func than Attr in futur?)
		if name == "value" || name == "checked" || name == "selected" {
			parent.Set(name.(string), value)
			return nil
		}

		parent.Call("setAttribute", name, value)
		OnCleanup(func() {
			parent.Call("removeAttribute", name)
		})

		return nil
	})
}

func BindAttr[T any](name string, value func() T) Node {
	return Bind(func() Node {
		return Attr(name, value())
	})
}
