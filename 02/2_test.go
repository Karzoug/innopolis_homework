package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countVotes(t *testing.T) {
	tests := []struct {
		names []string
		want  []Candidate
	}{
		{
			names: []string{"Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"},
			want:  []Candidate{{"Ann", 3}, {"Kate", 2}, {"Peter", 1}, {"Helen", 1}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, countVotes(tt.names))
	}
}
