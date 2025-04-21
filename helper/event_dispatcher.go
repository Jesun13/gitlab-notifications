package helper

import (
	"fmt"
)

type EventHandler interface {
	HandleEvent(event map[string]interface{}, topicKey string) error
}
type EventDispatcher struct {
	handlers map[string]EventHandler
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string]EventHandler),
	}
}

func (d *EventDispatcher) RegisterHandler(eventType string, handler EventHandler) {
	d.handlers[eventType] = handler
}

func (d *EventDispatcher) Dispatch(eventType string, event map[string]interface{}, topicKey string) error {
	handler, exists := d.handlers[eventType]
	if !exists {
		return fmt.Errorf("не найден обработчик для события: %s", eventType)
	}
	return handler.HandleEvent(event, topicKey)
}
