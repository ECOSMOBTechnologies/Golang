package utils

import (
	"practical/model"
	"sort"
	"strings"
)

//repetition this method its call repetition words find
func repetition(st string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

//GetRepetitionWords its return top ten words in books
func GetRepetitionWords(input string) []model.Books {

	var words = make(map[string]int)

	for index, element := range repetition(input) {
		words[index] = element
	}

	return sortMap(words)
}

//sortMap its call method sortMap
func sortMap(words map[string]int) []model.Books {
	n := map[int][]string{}
	var a []int
	for k, v := range words {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	var mapSort = make([]model.Books, 0)

	for _, k := range a {
		for _, s := range n[k] {
			if len(mapSort) == 10 {
				break
			}
			mapSort = append(mapSort, model.Books{Key: s, Val: k})
		}
	}
	return mapSort
}
