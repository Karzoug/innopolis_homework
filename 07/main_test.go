package main

import (
	"math/rand"
	"testing"
)

var arr []int

func init() {
	arr = make([]int, 1_000_000)
	for i := range arr {
		arr[i] = rand.Int()
	}
}

func BenchmarkMinEl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MinEl(arr)
	}
}

func BenchmarkMinEl2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MinEl2(arr)
	}
}

func BenchmarkMinEl3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MinEl3(arr)
	}
}
