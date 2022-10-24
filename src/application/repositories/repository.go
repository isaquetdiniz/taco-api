package repositories

type Repository[T comparable] interface {
	Save(*T) *T
	GetById(id string) *T
	Update(*T) *T
	DeleteById(id string)
}
