//go:build js && wasm

package components

import (
	"syscall/js"

	. "github.com/loom-go/loom"
	"github.com/loom-go/web/internal"
)

func Elem(name string, children ...Node) Node {
	return &elemNode{name: name, children: children}
}

type elemNode struct {
	name     string
	children []Node
}

func (n *elemNode) ID() string {
	return "web.Elem." + n.name
}

func (n *elemNode) Mount(slot *Slot) error {
	parent := slot.Parent().(js.Value)

	self := internal.Doc().Call("createElement", n.name)
	slot.SetSelf(self)

	parent.Call("appendChild", self)

	return n.Update(slot)
}

func (n *elemNode) Update(slot *Slot) error {
	return slot.RenderChildren(n.children...)
}

func (n *elemNode) Unmount(slot *Slot) error {
	if slot.Self() == nil {
		return nil
	}

	parent := slot.Parent().(js.Value)
	self := slot.Self().(js.Value)
	parent.Call("removeChild", self)

	return nil
}

func Div(children ...Node) Node    { return Elem("div", children...) }
func Span(children ...Node) Node   { return Elem("span", children...) }
func P(children ...Node) Node      { return Elem("p", children...) }
func H1(children ...Node) Node     { return Elem("h1", children...) }
func H2(children ...Node) Node     { return Elem("h2", children...) }
func H3(children ...Node) Node     { return Elem("h3", children...) }
func H4(children ...Node) Node     { return Elem("h4", children...) }
func H5(children ...Node) Node     { return Elem("h5", children...) }
func Ul(children ...Node) Node     { return Elem("ul", children...) }
func Ol(children ...Node) Node     { return Elem("ol", children...) }
func Li(children ...Node) Node     { return Elem("li", children...) }
func Form(children ...Node) Node   { return Elem("form", children...) }
func Button(children ...Node) Node { return Elem("button", children...) }
func Input(children ...Node) Node  { return Elem("input", children...) }
func Img(children ...Node) Node    { return Elem("img", children...) }
func A(children ...Node) Node      { return Elem("a", children...) }
