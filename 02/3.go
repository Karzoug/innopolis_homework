package main

import "fmt"

/* 3. У учеников старших классов прошел контрольный срез по нескольким предметам. Выведите данные в читаемом виде
в таблицу вида
_____________________________________
Student name  | Grade | Object    |  Result
____________________________________
Ann			  |     9 | Math	  |  4
Ann 		  |     9 | Biology   |  4
...

Вводные данные представлены в файле dz3.json
*/

type Student struct {
	ID    int
	Name  string
	Grade int
}

type Object struct {
	ID   int
	Name string
}

type Result struct {
	ObjectID  int `json:"object_id,omitempty"`
	StudentID int `json:"student_id,omitempty"`
	Result    int `json:"result,omitempty"`
}

type db struct {
	Students []Student
	Objects  []Object
	Results  []Result
}

func printTableStudentsResults(db db) {
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
	for _, r := range db.Results {
		student := studentsMap[r.StudentID]
		object := objectsMap[r.ObjectID]
		fmt.Printf("%-13s | %5d | %-9s | %d\n", student.Name, student.Grade, object.Name, r.Result)
	}
}
