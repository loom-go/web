//go:build js && wasm

package components

import (
	"fmt"
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
	. "github.com/AnatoleLucet/loom/components"
)

func Attr(name, value any) Node {
	return &attrNode{name: name, value: value}
}

type attrNode struct {
	name  any
	value any
}

func (n *attrNode) ID() string {
	return fmt.Sprintf("web.Attr.%v", n.name)
}

func (n *attrNode) Mount(slot *Slot) error {
	slot.SetNode(n)

	return n.Update(slot)
}

func (n *attrNode) Update(slot *Slot) error {
	parent := slot.Parent().(js.Value)
	slot.SetNode(n)

	if n.name == "value" || n.name == "checked" || n.name == "selected" {
		parent.Set(n.name.(string), n.value)
		return nil
	}

	parent.Call("setAttribute", n.name, n.value)
	return nil
}

func (n *attrNode) Unmount(slot *Slot) error {
	parent := slot.Parent().(js.Value)
	parent.Call("removeAttribute", n.name)
	return nil
}

func BindAttr[T any](name string, value func() T) Node {
	return Bind(func() Node {
		return Attr(name, value())
	})
}
