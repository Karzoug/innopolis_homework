package main

import (
	"errors"
	"log"
)

func main() {
}

func EvalSequence(matrix [][]int, userAnswer []int) int {
	if err := validation(matrix, userAnswer); err != nil {
		log.Println(err)
		return 0
	}

	maxGrade := calcMaxGrade(matrix)
	userGrade := calcUserGrade(matrix, userAnswer)

	return userGrade * 100 / maxGrade
}

func validation(matrix [][]int, userAnswer []int) error {
	// матрица должна быть квадратной
	length := len(matrix)
	for i := range matrix {
		if len(matrix[i]) != length {
			return errors.New("matrix is not square")
		}
	}
	// в графе не может быть петель (диагональ матрицы)
	for i := range matrix {
		if matrix[i][i] != 0 {
			return errors.New("matrix has loop")
		}
	}
	// ответы пользователя не должны выходить за диапазон матрицы
	for i := range userAnswer {
		if userAnswer[i] >= length || userAnswer[i] < 0 {
			return errors.New("user answer is out of range")
		}
	}
	// элементы в слайсе ответов пользователя должны быть уникальными
	set := make(map[int]struct{}, len(userAnswer))
	for i := range userAnswer {
		if _, exists := set[userAnswer[i]]; exists {
			return errors.New("user answer is not unique")
		}
		set[userAnswer[i]] = struct{}{}
	}

	return nil
}

func calcMaxGrade(matrix [][]int) int {
	var sum int
	for i := 0; i < len(matrix); i++ {
		path := dfsMaxPrice(i, matrix)
		if path > sum {
			sum = path
		}
	}
	return sum
}

func calcUserGrade(matrix [][]int, userAnswer []int) int {
	var sum int
	for i := 1; i < len(userAnswer); i++ {
		sum += matrix[userAnswer[i-1]][userAnswer[i]]
	}
	return sum
}

func dfsMaxPrice(startVertex int, adjMatrix [][]int) int {
	visited := make([]bool, len(adjMatrix))
	return dfsMaxPriceUtil(startVertex, 0, visited, adjMatrix)
}
func dfsMaxPriceUtil(vertex int, sum int, visited []bool, adjMatrix [][]int) int {
	visited[vertex] = true
	maxSum := sum
	for i := 0; i < len(adjMatrix); i++ {
		if adjMatrix[vertex][i] != 0 && !visited[i] {
			path := dfsMaxPriceUtil(i, sum+adjMatrix[vertex][i], visited, adjMatrix)
			maxSum = max(maxSum, path)
		}
	}
	return maxSum
}
