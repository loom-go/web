//go:build js && wasm

package main

import (
	"fmt"
	"time"

	"github.com/loom-go/loom"
	. "github.com/loom-go/loom/components"
	"github.com/loom-go/web"
	. "github.com/loom-go/web/components"
)

func App() loom.Node {
	count, setCount := Signal(0)

	go func() {
		for {
			time.Sleep(time.Second / 30)
			setCount(count() + 1)
		}
	}()

	return P(Text("Count: "), BindText(count))
}

func main() {
	app := web.NewApp()

	for err := range app.Run("#app", App) {
		fmt.Println("Error:", err) // todo: console error
	}
}
