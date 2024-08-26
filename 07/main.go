package main

/*
Распишите ваше решение в отчете. Отчет может быть представлен в любом текстовом формате:
	* комментарии в коде
	* Markdown
	* latex
	* docx
	* xlsx
	* фото листа читаемого качества, читаемого почерка
*/

/*
Представлен алгоритм для определения “зеркальности” матрицы размером n на n.
Вычислите его сложность, используя изученные формулы, в качестве размера входных данных примите n.

* Размер входных данных: n
* Основная операция алгоритма: сравнение
* Зависит ли число выполняемых основных операций только от размера входных данных: нет,
  количество операций сравнения будет зависеть не только от общего числа n элементов в массиве,
  но и от того, есть ли в массиве "незеркальные" элементы, и если есть,
  то на каких позициях они расположены.
* Рекурентное уравнение, выражающее количество выполняемых основных операций:
	M(n) = ∑(0, n-2) ∑(0, n-1) 1 = ∑(0, n-1) (n) = n * ∑(0, n-1) 1 = n * n


	===>

	M(n) = n^2 и O(n^2)


*/

func IsReflectMatrix(a [][]int) bool {
	n := len(a)
	if n == 0 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if len(a[i]) != n {
			return false
		}
		for j := 0; j < n; j++ {
			if a[i][j] != a[j][i] {
				return false
			}
		}
	}
	return true
}

/*
Представлен рекурсивный алгоритм для поиска наименьшего элемента слайса.
Вычислите его сложность, используя изученные формулы.

* Размер входных данных: n
* Основная операция алгоритма: сравнение
* Зависит ли число выполняемых основных операций от размера входных данных: да
* Рекурентное уравнение, выражающее количество выполняемых основных операций:
	M(n) = 1 + M(n-1)
	M(1) = 0
	M(0) = 0

	===>

	M(n) = n и O(n)
*/

func MinEl(a []int) int {
	// только для первичной проверки
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	t := MinEl(a[:len(a)-1])
	if t <= a[len(a)-1] {
		return t
	}
	return a[len(a)-1]
}

/*
Представлен другой рекурсивный алгоритм для поиска наименьшего элемента слайса.
Вычислите его сложность, используя изученные формулы.

* Размер входных данных: n
* Основная операция алгоритма: сравнение
* Зависит ли число выполняемых основных операций от размера входных данных: да
* Рекурентное уравнение, выражающее количество выполняемых основных операций:
	M(n) = 1 + M(⌊n/2⌋) + M(⌈n/2⌉)
	M(1) = 0

	если принять, что n равно 2^k, то

	M(2^k) = 1 + 2 * M(2^(k-1))
	M(2^0) = 0

	===>

	M(n) = 2^k = 2 ^ log_2(n) = n и O(n)


Сравните эффективность алгоритма с предыдущим вариантом.
Попробуйте увеличить скорость выполнения функции, используя инструменты Го.
Как изменится сложность алгоритма при увеличении скорости?

задача решается оптимально простым проходом цикла, т.е. без использования рекурсии,
сложность алгоритма будет также O(n) - меньше она быть не может, так как мы должны просмотреть все элементы массива
см. MinEl3 и main_test.go

*/

func MinEl2(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	t1 := MinEl2(a[:len(a)/2])
	t2 := MinEl2(a[len(a)/2:])
	if t1 <= t2 {
		return t1
	}
	return t2
}

func MinEl3(a []int) int {
	if len(a) == 0 {
		return 0
	}

	minEl := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < minEl {
			minEl = a[i]
		}
	}
	return minEl
}
