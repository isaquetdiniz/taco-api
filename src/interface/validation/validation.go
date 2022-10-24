package validation

type Validator interface {
	Validate(ops ...interface{}) error
}
