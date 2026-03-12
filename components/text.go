//go:build js && wasm

package components

import (
	"syscall/js"

	. "github.com/loom-go/loom"
	. "github.com/loom-go/loom/components"
	"github.com/loom-go/web/internal"
)

func Text(content any) Node {
	return &textNode{content: content}
}

type textNode struct {
	content any
}

func (n *textNode) ID() string {
	return "web.Text"
}

func (n *textNode) Mount(slot *Slot) error {
	parent := slot.Parent().(*js.Value)

	self := internal.Doc().Call("createTextNode", n.content)
	slot.SetSelf(&self)

	parent.Call("appendChild", self)

	return nil
}

func (n *textNode) Update(slot *Slot) error {
	self := slot.Self().(*js.Value)
	self.Set("nodeValue", n.content)

	return nil
}

func (n *textNode) Unmount(slot *Slot) error {
	if slot.Self() == nil {
		return nil
	}

	parent := slot.Parent().(*js.Value)
	self := slot.Self().(*js.Value)
	parent.Call("removeChild", self)

	return nil
}

func BindText[T any](value func() T) Node {
	return Bind(func() Node {
		return Text(value())
	})
}
