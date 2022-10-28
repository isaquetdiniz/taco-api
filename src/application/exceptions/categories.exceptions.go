package exceptions

import "errors"

var (
	CategoryAlreadyExists = errors.New("Category already exits")
	CategoryNotFound      = errors.New("Category not found")
)
