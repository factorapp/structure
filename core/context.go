package core

import (
	dom "github.com/gowasm/go-js-dom"
)

// Context gets passed into event handlers. It's super helpful for rad things!
type Context interface {
	dom.Event
}

type contextEventWrapper struct {
	dom.Event
}
