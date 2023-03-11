package aggregate

type EventRecorder struct {
	events []interface{}
}

func (r *EventRecorder) AddEvent(event interface{}) {
	r.events = append(r.events, event)
}

func (r *EventRecorder) Events() []interface{} {
	return r.events
}

func (r *EventRecorder) ClearEvents() {
	r.events = nil
}
