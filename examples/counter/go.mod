module myapp

go 1.25.5

replace github.com/loom-go/web => ../..

require (
	github.com/loom-go/loom v0.2.0
	github.com/loom-go/web v0.0.0-00010101000000-000000000000
)

require (
	github.com/AnatoleLucet/sig v0.0.0-20260308162001-17251018b48a // indirect
	github.com/petermattis/goid v0.0.0-20251121121749-a11dd1a45f9a // indirect
)
