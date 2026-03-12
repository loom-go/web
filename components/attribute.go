//go:build js && wasm

package components

import (
	"reflect"
	"syscall/js"

	"github.com/loom-go/loom/components"
)

type Attr = Attribute

type Attribute map[string]any

func (a Attribute) Apply(parent any) (func() error, error) {
	p := parent.(*js.Value)

	var removers []func() error
	for name, value := range a {
		removers = append(removers, a.applyAttr(p, name, value))
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

func (a Attribute) applyAttr(parent *js.Value, name string, value any) func() error {
	val, vok := unwrapAccessor[any](value)
	if !vok {
		return func() error { return nil }
	}

	if name == "value" || name == "checked" || name == "selected" {
		parent.Set(name, val)
		return func() error { return nil }
	}

	parent.Call("setAttribute", name, val)

	return func() error {
		parent.Call("removeAttribute", name)
		return nil
	}
}

// little helper to unwrap a `T | func() T`. mainly used in Appliers
func unwrapAccessor[V any](value any) (v V, vok bool) {
	if sig, ok := value.(func() V); ok {
		value = sig()
	} else if fn, ok := value.(func() any); ok {
		value = fn()
	} else if fn, ok := value.(components.Accessor[V]); ok {
		value = fn()
	} else if fn, ok := value.(components.Accessor[any]); ok {
		value = fn()
	} else {
		// fallback to reflect for calling the accessor
		rv := reflect.ValueOf(value)
		isFunc := rv.Kind() == reflect.Func
		if isFunc && rv.Type().NumIn() == 0 && rv.Type().NumOut() == 1 {
			value = rv.Call(nil)[0].Interface()
		}
	}

	if val, ok := value.(V); ok {
		v = val
		vok = true
	} else {
		// fallback to reflect for converting the value
		rv := reflect.ValueOf(value)
		target := reflect.TypeFor[V]()
		if rv.IsValid() && rv.Type().ConvertibleTo(target) {
			v = rv.Convert(target).Interface().(V)
			vok = true
		}
	}

	return
}
