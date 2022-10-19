package entities

import (
	"fmt"
	"time"
)

type Entity struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e Entity) String() string {
	return fmt.Sprintf("Categoria: %s [%s]", e.Name, e.Id)
}
