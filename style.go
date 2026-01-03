//go:build js && wasm

package web

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
)

func Style(styles map[string]string) Node {
	return NodeFunc(func(ctx *RenderContext) error {
		parent := ctx.Get("parent").(js.Value)
		style := parent.Get("style")

		for key, value := range styles {
			style.Set(key, value)
		}

		return nil
	})
}
