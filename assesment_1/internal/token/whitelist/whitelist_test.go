package whitelist

import (
	"testing"

	"assesment_1/internal/repo/mxmap"

	"github.com/stretchr/testify/assert"
)

func Test_tokenValidator_Validate(t *testing.T) {
	const (
		token1 = "hDFJ!4&5MVg*bTDX"
		token2 = "hlDHur3wr$%RHRS2"
		token3 = "hDFJ!4&5MVg*bTDX2"
	)

	set := mxmap.New[string, struct{}]()

	tokenValidator := New(set)

	tokenValidator.Add(token1)
	tokenValidator.Add(token2)

	assert.True(t, tokenValidator.Validate(token1))
	assert.True(t, tokenValidator.Validate(token2))

	assert.False(t, tokenValidator.Validate(token3))

	tokenValidator.Remove(token1)
	assert.False(t, tokenValidator.Validate(token1))
}
