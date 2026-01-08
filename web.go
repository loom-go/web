//go:build js && wasm

package web

import (
	"errors"

	"github.com/AnatoleLucet/loom"
	"github.com/AnatoleLucet/loom-web/internal"
	"github.com/AnatoleLucet/loom/signals"
)

type App struct {
	running bool

	owner *signals.Owner
}

func NewApp() *App {
	return &App{
		running: false,
		owner:   signals.NewOwner(),
	}
}

func (a *App) Run(parent string, fn func() loom.Node) <-chan any {
	errc := make(chan any, 1)
	if a.running {
		errc <- errors.New("app is already running")
		return errc
	}

	a.owner.OnError(func(err any) {
		errc <- err
	})

	err := a.Render(parent, fn)
	if err != nil {
		errc <- err
		return errc
	}

	a.running = true
	return errc
}

func (a *App) Render(parent string, fn func() loom.Node) error {
	err := a.owner.Run(func() error {
		container := internal.Doc().Call("querySelector", parent)
		return loom.Render(container, fn())
	})

	if err != nil {
		a.owner.Dispose()
		return err
	}

	return nil
}
