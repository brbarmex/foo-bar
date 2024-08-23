package main

import (
	"log"
)

func main() {

	var collections = []int{1, 2, 3, 4, 5, 6}

	var index_five int

	for keyCollevtion, valueCollection := range collections {
		if valueCollection == 5 {
			index_five = keyCollevtion
		}
	}

	log.Printf("the index %d contains the value 5", index_five)

	value := Slice.FilteredIter(collections, func(i int) bool {
		return i == 5
	})

	log.Print(value)

}

type Slice []int

func (s Slice) FilteredIter(predicate func(i int) bool) int {
	for i := range s {
		if predicate(s[i]) {
			return s[i]
		}
	}

	return 0
}
