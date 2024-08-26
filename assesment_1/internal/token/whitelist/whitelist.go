package whitelist

import (
	"assesment_1/internal/repo"
	"assesment_1/internal/token"
)

var _ token.Validator = (*tokenValidator)(nil)

type tokenValidator struct {
	set repo.CMap[string, struct{}]
}

func New(set repo.CMap[string, struct{}]) *tokenValidator {
	return &tokenValidator{
		set: set,
	}
}

func (t *tokenValidator) Validate(token string) bool {
	_, ok := t.set.Load(token)
	return ok
}

func (t *tokenValidator) Add(token string) {
	t.set.Store(token, struct{}{})
}

func (t *tokenValidator) Remove(token string) {
	t.set.Delete(token)
}
