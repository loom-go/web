//go:build js && wasm

package components

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
)

type Style map[string]any

func (s Style) Render(ctx *RenderContext) error {
	parent := ctx.Get("parent").(js.Value)
	style := parent.Get("style")

	for key, value := range s {
		style.Set(key, value)
	}

	return nil
}
