//go:build js && wasm

package web

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
	"github.com/AnatoleLucet/sig"
)

func Render(parent string, node Node) error {
	// todo: return root owner here with error handling to avoid panics
	return render(NewRenderContext(), doc().Call("querySelector", parent), node)
}

func render(ctx *RenderContext, parent js.Value, children ...Node) error {
	for _, child := range children {
		childCtx := ctx.Clone()
		childCtx.Set("parent", parent)

		// todo: render shouldn't have to worry about ownership, that should be implicit to the node somehow
		err := sig.NewOwner().Run(func() error {
			return child.Render(childCtx)
		})
		if err != nil {
			return err
		}
	}

	return nil
}
