package event

import (
	"github.com/google/uuid"
)

type DomainEvent interface {
	ID() uuid.UUID
	Type() string
	AggregateID() string
	AggregateType() string
}

type BaseEvent struct {
	id          uuid.UUID
	aggregateID string
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		id:          uuid.Must(uuid.NewV7()),
		aggregateID: aggregateID,
	}
}

func (e BaseEvent) ID() uuid.UUID       { return e.id }
func (e BaseEvent) AggregateID() string { return e.aggregateID }
