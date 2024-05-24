package main

import (
	"fmt"
	"slices"

	"golang.org/x/exp/maps"
)

/* 6. Перепишите задачу #4 с использованием функций высшего порядка, изученных на лекции.
Желательно реализуйте эти функции самостоятельно.


*/

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[T1, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
	r := init
	for _, v := range s {
		r = f(v, r)
	}
	return r
}

func printTableMeanScoreByObjectAndGradeFunctional(db db) {
	if len(db.Results) == 0 {
		fmt.Println("No data")
		return
	}

	studentsCache := NewCache[int, Student]()
	for _, s := range db.Students {
		studentsCache.Set(s.ID, s)
	}

	for _, object := range db.Objects {
		gradesSet := make(map[int]struct{})
		filteredByObject := Filter(db.Results, func(r Result) bool {
			if r.ObjectID == object.ID {
				student, _ := studentsCache.Get(r.StudentID)
				gradesSet[student.Grade] = struct{}{}
				return true
			}
			return false
		})
		if len(filteredByObject) == 0 {
			continue
		}

		printHorizontalRule()
		printObjectHeader(object.Name)
		printHorizontalRule()
		meanByObject := Reduce(filteredByObject, 0.0, func(r Result, acc float64) float64 {
			return acc + float64(r.Result)
		}) / float64(len(filteredByObject))

		gradesSlice := maps.Keys(gradesSet)
		slices.Sort(gradesSlice)

		for _, grade := range gradesSlice {
			filteredByObjectAndGrade := Filter(filteredByObject, func(r Result) bool {
				student, _ := studentsCache.Get(r.StudentID)
				return student.Grade == grade
			})
			meanByObjectAndGrade := Reduce(filteredByObjectAndGrade, 0.0, func(r Result, acc float64) float64 {
				return acc + float64(r.Result)
			}) / float64(len(filteredByObjectAndGrade))
			printMeanResultInGrade(grade, meanByObjectAndGrade)
		}
		printHorizontalRule()
		printMeanResultInObject(meanByObject)
	}

	printHorizontalRule()
}
