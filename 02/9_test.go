package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumbersSum(t *testing.T) {
	tests := []struct {
		args Numbers[uint]
		want uint
	}{
		{
			args: []uint{1, 2, 3},
			want: 6,
		},
		{
			args: []uint{1, 2, 4},
			want: 7,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.args.Sum())
	}
}

func TestNumbersProduct(t *testing.T) {
	tests := []struct {
		args Numbers[uint]
		want uint
	}{
		{
			args: []uint{1, 2, 3},
			want: 6,
		},
		{
			args: []uint{1, 2, 4},
			want: 8,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.args.Product())
	}
}

func TestNumbersCompare(t *testing.T) {
	tests := []struct {
		args [2]Numbers[uint]
		want bool
	}{
		{
			args: [2]Numbers[uint]{
				[]uint{1, 2, 3},
				[]uint{1, 2, 3},
			},
			want: true,
		},
		{
			args: [2]Numbers[uint]{
				[]uint{1, 2, 3},
				[]uint{1, 2, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.args[0].Compare(tt.args[1]))
	}
}

func TestNumbersFind(t *testing.T) {
	type args struct {
		v   uint
		arr Numbers[uint]
	}
	tests := []struct {
		args      args
		wantIndex int
		wantFound bool
	}{
		{
			args: args{
				v:   1,
				arr: []uint{1, 2, 3},
			},
			wantIndex: 0,
			wantFound: true,
		},
		{
			args: args{
				v:   2,
				arr: []uint{1, 2, 3},
			},
			wantIndex: 1,
			wantFound: true,
		},
		{
			args: args{
				v:   8,
				arr: []uint{1, 2, 3},
			},
			wantIndex: 3,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		gotIndex, gotFound := tt.args.arr.Find(tt.args.v)
		assert.Equal(t, tt.wantIndex, gotIndex)
		assert.Equal(t, tt.wantFound, gotFound)
	}
}

func TestNumbersDelete(t *testing.T) {
	type args struct {
		v   uint
		arr Numbers[uint]
	}
	tests := []struct {
		args args
		want []uint
	}{
		{
			args: args{
				v:   1,
				arr: []uint{1, 2, 3},
			},
			want: []uint{2, 3},
		},
		{
			args: args{
				v:   5,
				arr: []uint{1, 2, 3},
			},
			want: []uint{1, 2, 3},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.args.arr.Delete(tt.args.v))
	}
}

func TestNumbersDeleteByIndex(t *testing.T) {
	type args struct {
		index int
		arr   Numbers[uint]
	}
	tests := []struct {
		args args
		want []uint
	}{
		{
			args: args{
				index: 1,
				arr:   []uint{1, 2, 3},
			},
			want: []uint{1, 3},
		},
		{
			args: args{
				index: 0,
				arr:   []uint{1, 2, 3},
			},
			want: []uint{2, 3},
		},
		{
			args: args{
				index: 2,
				arr:   []uint{1, 2, 3},
			},
			want: []uint{1, 2},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.args.arr.DeleteByIndex(tt.args.index))
	}
}
