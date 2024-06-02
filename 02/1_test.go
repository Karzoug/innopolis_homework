package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersection(t *testing.T) {
	tests := []struct {
		arr  [][]int
		want []int
	}{
		{
			arr: [][]int{
				{1, 2, 3, 2},
				{3, 2},
			},
			want: []int{2, 3},
		},
		{
			arr: [][]int{
				{1, 2, 3, 2},
			},
			want: []int{1, 2, 3},
		},
		{
			arr: [][]int{
				{1, 2, 3, 2},
				{3, 2},
				{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, intersection(tt.arr...))
	}
}
