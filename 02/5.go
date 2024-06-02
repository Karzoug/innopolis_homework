package main

import "fmt"

/* 5. Перепишите задачу #3 с использованием структуры-дженерик Cache, изученной на семинаре.
Храните в кеше таблицы студентов и предметов.
*/

type Cache[K comparable, V any] struct {
	m map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		m: make(map[K]V),
	}
}
func (c *Cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}
func (c *Cache[K, V]) Get(key K) (V, bool) {
	k, ok := c.m[key]
	return k, ok
}

func (c *Cache[K, V]) Delete(key K) {
	delete(c.m, key)
}

func (c *Cache[K, V]) Clear() {
	clear(c.m)
}

func printTableStudentsResultsGeneric(db db) {
	studentsCache := NewCache[int, Student]()
	for _, s := range db.Students {
		studentsCache.Set(s.ID, s)
	}
	objectsCache := NewCache[int, Object]()
	for _, o := range db.Objects {
		objectsCache.Set(o.ID, o)
	}

	fmt.Print(`_____________________________________
Student name  | Grade | Object    |  Result
____________________________________
`)
	for _, r := range db.Results {
		student, _ := studentsCache.Get(r.StudentID)
		object, _ := objectsCache.Get(r.ObjectID)
		fmt.Printf("%-13s | %5d | %-9s | %d\n", student.Name, student.Grade, object.Name, r.Result)
	}
}
