//go:build js && wasm

package components

import (
	"syscall/js"

	. "github.com/AnatoleLucet/loom"
)

type Style map[string]any

func (s Style) ID() string {
	return "web.Style"
}

func (s Style) Mount(slot *Slot) error {
	return s.Update(slot)
}

func (s Style) Update(slot *Slot) error {
	parent := slot.Parent().(js.Value)
	style := parent.Get("style")
	slot.SetNode(s)

	for key, value := range s {
		style.Set(key, value)
	}

	return nil
}

func (s Style) Unmount(slot *Slot) error {
	parent := slot.Parent().(js.Value)
	style := parent.Get("style")

	for key := range s {
		style.Set(key, "")
	}

	return nil
}
