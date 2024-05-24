package main

import (
	"slices"

	"golang.org/x/exp/constraints"
)

/* 9. Реализуйте тип-дженерик Numbers, который является слайсом численных типов.
Реализуйте следующие методы для этого типа:
* суммирование всех элементов,
* произведение всех элементов,
* сравнение с другим слайсом на равность,
* проверка аргумента, является ли он элементом массива, если да - вывести индекс первого найденного элемента,
* удаление элемента массива по значению,
* удаление элемента массива по индексу.
*/

type Numbers[T constraints.Integer] []T

func (n Numbers[T]) Sum() T {
	sum := T(0)
	for _, v := range n {
		sum += v
	}
	return sum
}

func (n Numbers[T]) Product() T {
	product := T(1)
	for _, v := range n {
		product *= v
	}
	return product
}

func (n Numbers[T]) Compare(other Numbers[T]) bool {
	if len(n) != len(other) {
		return false
	}
	for i := range n {
		if n[i] != other[i] {
			return false
		}
	}
	return true
}

func (n Numbers[T]) Find(v T) (int, bool) {
	for i := range n {
		if v == n[i] {
			return i, true
		}
	}
	return len(n), false
}

func (n Numbers[T]) Delete(v T) []T {
	return slices.DeleteFunc(n,
		func(x T) bool {
			return x == v
		})
}

func (n Numbers[T]) DeleteByIndex(i int) []T {
	return slices.Delete(n, i, i+1)
}
