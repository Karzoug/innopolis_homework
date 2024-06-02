package main

import (
	"slices"

	"golang.org/x/exp/maps"
)

/*
 1. Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.

Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
Примеры:
1. Если на вход подается только 1 слайс [1, 2, 3, 2], результатом должен быть слайс [1, 2, 3].
2. Вход: 2 слайса [1, 2, 3, 2] и [3, 2], результат - [2, 3].
3. Вход: 3 слайса [1, 2, 3, 2], [3, 2] и [], результат - [].
*/
func intersection(arr ...[]int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	indexArrMinLen := 0
	minCount := len(arr[0])
	for i, v := range arr {
		if len(v) < minCount {
			minCount = len(v)
			indexArrMinLen = i
		}
	}

	seen := make(map[int]bool)
	for i := range arr[indexArrMinLen] {
		seen[arr[indexArrMinLen][i]] = false
	}

	for i := range arr {
		for j := range arr[i] {
			if _, found := seen[arr[i][j]]; found {
				seen[arr[i][j]] = true
			}
		}
		for k, v := range seen {
			if v {
				seen[k] = false
			} else {
				delete(seen, k)
			}
		}
	}

	res := maps.Keys(seen)

	slices.Sort(res)

	return res
}
