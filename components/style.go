//go:build js && wasm

package components

import (
	"syscall/js"
)

// todo: put static styles in a <style> in the head instead of applying them inline
// just apply inline the reactive styles.

type Style map[string]any

func (s Style) Apply(parent any) (func() error, error) {
	p := parent.(*js.Value)

	var removers []func() error
	for key, value := range s {
		removers = append(removers, s.applyStyle(p, key, value))
	}

	return func() error {
		for _, remove := range removers {
			if err := remove(); err != nil {
				return err
			}
		}

		return nil
	}, nil
}

func (s Style) applyStyle(parent *js.Value, key string, value any) func() error {
	val, vok := unwrapAccessor[any](value)
	if !vok {
		return func() error { return nil }
	}

	parent.Get("style").Set(key, val)

	return func() error {
		parent.Get("style").Set(key, "")
		return nil
	}
}
