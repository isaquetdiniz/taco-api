package dtos

import (
	"time"
)

type EntityDTO struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
