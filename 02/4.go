package main

import (
	"fmt"
	"slices"
)

/* 4. Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам в виде:
________________
Math	 | Mean
________________
 9 grade | 4.5
10 grade | 5
11 grade | 3.5
________________
mean     | 4		- среднее значение среди всех учеников
________________
________________
Biology	 | Mean
________________
...
Вводные данные представлены в файле dz3.json
*/

func printTableMeanScoreByObjectAndGrade(db db) {
	studentsMap := make(map[int]Student)
	for _, s := range db.Students {
		studentsMap[s.ID] = s
	}

	objectsMap := make(map[int]Object)
	for _, o := range db.Objects {
		objectsMap[o.ID] = o
	}

	slices.SortFunc(db.Results, func(a, b Result) int {
		if diff := a.ObjectID - b.ObjectID; diff != 0 {
			return diff
		}
		return studentsMap[a.StudentID].Grade - studentsMap[b.StudentID].Grade
	})

	if len(db.Results) == 0 {
		fmt.Println("No data")
		return
	}
	prevObjectID := db.Results[0].ObjectID - 1 // special value for first iteration
	prevGrade := studentsMap[db.Results[0].StudentID].Grade
	var (
		sumByGrade, countByGrade   float64
		sumByObject, countByObject float64
	)

	for _, r := range db.Results {
		if prevGrade != studentsMap[r.StudentID].Grade ||
			(r.ObjectID != prevObjectID && countByGrade != 0) {
			printMeanResultInGrade(prevGrade, sumByGrade/countByGrade)
			prevGrade = studentsMap[r.StudentID].Grade
			sumByGrade = 0
			countByGrade = 0
		}
		if r.ObjectID != prevObjectID {
			if countByObject != 0 {
				printHorizontalRule()
				printMeanResultInObject(sumByObject / countByObject)
			}
			printHorizontalRule()
			printObjectHeader(objectsMap[r.ObjectID].Name)
			printHorizontalRule()
			prevObjectID = r.ObjectID
			sumByObject = 0
			countByObject = 0
		}
		sumByGrade += float64(r.Result)
		countByGrade++

		sumByObject += float64(r.Result)
		countByObject++
	}

	printMeanResultInGrade(prevGrade, sumByGrade/countByGrade)
	printHorizontalRule()
	if countByObject != 0 {
		printMeanResultInObject(sumByObject / countByObject)
	}
}

func printObjectHeader(objectName string) {
	fmt.Printf("%-9s | Mean\n", objectName)
}

func printHorizontalRule() {
	fmt.Println("________________")
}

func printMeanResultInGrade(grade int, mean float64) {
	fmt.Printf("%-3d grade | %.1f\n", grade, mean)
}

func printMeanResultInObject(mean float64) {
	fmt.Printf("mean      | %.1f\n", mean)
}
