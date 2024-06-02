package main

import "fmt"

/* 7. Выведите в консоль круглых отличников из числа студентов, используя функцию Filter.
Вывод реализуйте как в задаче #3.

_____________________________________
Student name  | Grade | Object    |  Result
____________________________________
Ann			  |     9 | Math	  |  4
Ann 		  |     9 | Biology   |  4
...

*/

func printTableAStudentsResults(db db) {
	studentsMap := make(map[int]Student)
	for _, s := range db.Students {
		studentsMap[s.ID] = s
	}

	objectsMap := make(map[int]Object)
	for _, o := range db.Objects {
		objectsMap[o.ID] = o
	}

	fmt.Print(`_____________________________________
Student name  | Grade | Object    |  Result
____________________________________
`)
	for id := range db.Students {
		aFlag := true
		filteredByStudent := Filter(db.Results, func(r Result) bool {
			if r.StudentID == id {
				if r.Result != 5 {
					aFlag = false
				}
				return true
			}
			return false
		})
		if aFlag {
			for _, r := range filteredByStudent {
				object := objectsMap[r.ObjectID]
				student := studentsMap[r.StudentID]
				fmt.Printf("%-13s | %5d | %-9s | %d\n", student.Name, student.Grade, object.Name, r.Result)
			}
		}
	}
}
