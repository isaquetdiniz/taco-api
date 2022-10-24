package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Option func(*Entity)

func WithID(ID string) Option {
	return func(e *Entity) {
		e.ID = ID
	}
}

func WithCreatedAt(CreatedAt time.Time) Option {
	return func(e *Entity) {
		e.CreatedAt = CreatedAt
	}
}

func WithUpdatedAt(UpdatedAt time.Time) Option {
	return func(e *Entity) {
		e.UpdatedAt = UpdatedAt
	}
}

func NewEntity(name string, options ...Option) *Entity {
	e := &Entity{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	for _, option := range options {
		option(e)
	}

	return e
}

func (e Entity) String() string {
	return fmt.Sprintf("[ID]: %s\n[Name]: %s\n[CreatedAt]: %s\n[UpdatedAt]: %s\n", e.ID, e.Name, e.CreatedAt, e.UpdatedAt)
}
