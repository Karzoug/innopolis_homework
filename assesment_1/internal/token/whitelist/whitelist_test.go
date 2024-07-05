package whitelist

import (
	"testing"

	"assesment_1/internal/repo/mxmap"

	"github.com/stretchr/testify/assert"
)

func Test_tokenValidator_Validate(t *testing.T) {
	set := mxmap.New[string, struct{}]()

	tokenValidator := New(set)

	tokenValidator.Add("hDFJ!4&5MVg*bTDX")
	tokenValidator.Add("hlDHur3wr$%RHRS2")

	assert.True(t, tokenValidator.Validate("hDFJ!4&5MVg*bTDX"))
	assert.True(t, tokenValidator.Validate("hlDHur3wr$%RHRS2"))

	assert.False(t, tokenValidator.Validate("hDFJ!4&5MVg*bTDX2"))

	tokenValidator.Remove("hDFJ!4&5MVg*bTDX")
	assert.False(t, tokenValidator.Validate("hDFJ!4&5MVg*bTDX"))
}
