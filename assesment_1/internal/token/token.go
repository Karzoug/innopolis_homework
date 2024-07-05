package token

type Validator interface {
	Validate(token string) bool
}
