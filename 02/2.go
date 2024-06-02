package main

import "slices"

/* 2. Подсчет голосов.
Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
Пример.
Вход: ["Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"]
Вывод: [{Ann, 3}, {Kate, 2}, {Peter, 1}, {Helen, 1}]
*/

type Candidate struct {
	Name  string
	Votes int
}

func countVotes(names []string) []Candidate {
	m := make(map[string]int)
	for i := range names {
		m[names[i]]++
	}

	res := make([]Candidate, len(m))
	var i int
	for k, v := range m {
		res[i] = Candidate{Name: k, Votes: v}
		i++
	}
	slices.SortFunc(res, func(a, b Candidate) int {
		return b.Votes - a.Votes
	})
	return res
}
