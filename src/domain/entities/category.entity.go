package entities

type Category struct {
	*Entity
}

func NewCategory(name string, options ...Option) *Category {
	c := &Category{
		Entity: NewEntity(name, options...),
	}

	return c
}
