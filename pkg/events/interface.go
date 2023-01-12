package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetPayload() interface{}
	GetDateTime() time.Time
	SetPayload(payload interface{})
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Dispatch(event EventInterface) error
	Register(eventName string, handler EventHandlerInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}
