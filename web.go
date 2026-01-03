//go:build js && wasm

package web

import (
	"github.com/AnatoleLucet/loom"
	"github.com/AnatoleLucet/loom-web/internal"
	"github.com/AnatoleLucet/loom/signals"
)

func Render(parent string, node loom.Node) (*signals.Owner, error) {
	o := signals.NewOwner()
	o.OnError(func(err any) {
		internal.ConsoleError("Recovered from panic:", err)
	})

	err := o.Run(func() error {
		return internal.RenderNodes(loom.NewRenderContext(), internal.Doc().Call("querySelector", parent), node)
	})
	if err != nil {
		o.Dispose()
		return nil, err
	}

	return o, nil
}
