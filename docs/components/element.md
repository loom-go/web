---
weight: 1
title: Elem()
---

```go {style=tokyonight-moon}
func Elem(name string, children ...Node) Node

func Br() Node                     { return Elem("br") }
func Div(children ...Node) Node    { return Elem("div", children...) }
func Span(children ...Node) Node   { return Elem("span", children...) }
func P(children ...Node) Node      { return Elem("p", children...) }
func S(children ...Node) Node      { return Elem("s", children...) }
func B(children ...Node) Node      { return Elem("b", children...) }
func I(children ...Node) Node      { return Elem("i", children...) }
func U(children ...Node) Node      { return Elem("u", children...) }
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
```
