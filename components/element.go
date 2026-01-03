//go:build js && wasm

package components

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
	"github.com/AnatoleLucet/loom-web/internal"
	. "github.com/AnatoleLucet/loom/signals"
)

func Elem(name string, children ...Node) Node {
	return NodeFunc(func(ctx *RenderContext) error {
		elem := internal.Doc().Call("createElement", name)

		err := internal.RenderNodes(ctx, elem, children...)
		if err != nil {
			return err
		}

		parent := ctx.Get("parent").(js.Value)
		parent.Call("appendChild", elem)

		OnCleanup(func() {
			parent.Call("removeChild", elem)
		})

		return nil
	})
}

func Input(typ string, attrs ...Node) Node {
	attrs = append([]Node{Attr("type", typ)}, attrs...)
	return Elem("input", attrs...)
}
func BindInput(typ func() string, attrs ...Node) Node {
	attrs = append([]Node{BindAttr("type", typ)}, attrs...)
	return Elem("input", attrs...)
}
func Img(src string, attrs ...Node) Node {
	attrs = append([]Node{Attr("src", src)}, attrs...)
	return Elem("img", attrs...)
}
func BindImg(src func() string, attrs ...Node) Node {
	attrs = append([]Node{BindAttr("src", src)}, attrs...)
	return Elem("img", attrs...)
}
func A(href string, children ...Node) Node {
	attrs := []Node{Attr("href", href)}
	attrs = append(attrs, children...)
	return Elem("a", attrs...)
}
func BindA(href func() string, children ...Node) Node {
	attrs := []Node{BindAttr("href", href)}
	attrs = append(attrs, children...)
	return Elem("a", attrs...)
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
