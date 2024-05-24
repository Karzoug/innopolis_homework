package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqualArrays(t *testing.T) {
	type args struct {
		a, b []uint
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				a: []uint{1, 2, 3},
				b: []uint{1, 2, 3},
			},
			want: true,
		},
		{
			args: args{
				a: []uint{1, 2, 3},
				b: []uint{1, 2, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, IsEqualArrays(tt.args.a, tt.args.b))
	}
}
