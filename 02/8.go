package main

import "slices"

/* 8. Напишите функцию-дженерик IsEqualArrays для comparable типов, которая сравнивает два неотсортированных массива.
Функция выдает булевое значение как результат. true - если массивы равны, false - если нет.
Массивы считаются равными, если в элемент из первого массива существует в другом, и наоборот.
Вне зависимости от расположения.
*/

func IsEqualArrays[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		// primitive logic
		// we don't use sort func because type is comparable not ordered
		if !slices.Contains(b, v) {
			return false
		}
	}
	return true
}
