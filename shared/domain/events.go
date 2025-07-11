package domain

import (
	"sync"
)

const (
	ImageUploaded = iota
	ImageTransformed
)

type Event interface{}

type Handler func(Event)

type EventBus struct {
	handlers map[int][]Handler
	sync.RWMutex
}

func (eb *EventBus) Subscribe(eventName int, handler Handler) {
	eb.Lock()
	defer eb.Unlock()
	eb.handlers[eventName] = append(eb.handlers[eventName], handler)
}

func (eb *EventBus) Publish(eventName int, event Event) {
	eb.RLock()
	defer eb.RUnlock()
	if handlers, ok := eb.handlers[eventName]; ok {
		for _, handler := range handlers {
			go handler(event)
		}
	}
}

var instance *EventBus

func New() *EventBus {
	if instance != nil {
		return instance
	}

	instance = &EventBus{
		handlers: make(map[int][]Handler),
	}

	return instance
}
