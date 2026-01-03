//go:build js && wasm

package internal

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
)

func RenderNodes(ctx *RenderContext, parent js.Value, children ...Node) error {
	for _, child := range children {
		childCtx := ctx.Clone()
		childCtx.Set("parent", parent)

		err := child.Render(childCtx)
		if err != nil {
			return err
		}
	}

	return nil
}
