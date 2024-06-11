package generics

import (
	"log"

	"golang.org/x/exp/constraints"
)

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float64
}

func Includes[T comparable](list []T, value T) (T, bool) {
	for _, v := range list {
		if v == value {
			return v, true
		}
	}

	var zero T

	return zero, false
}

func Filter[T constraints.Ordered](list []T, fn func(T) bool) []T {
	var filtered []T

	for _, v := range list {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func Constrains() {
	strings := []string{"pear", "apple", "pineapple", "orange", "banana"}
	numbers := []int{3, 8, 1, 7, 5, 4, 7, 2, 9, 10}

	value, ok := Includes(strings, "pear")

	if ok {
		log.Println(value, ok)
	} else {
		log.Println(value, ok)
	}

	val, ok := Includes(numbers, 50)

	if ok {
		log.Println(val, ok)
	} else {
		log.Println(val, ok)
	}

	log.Println(Filter(numbers, func(i int) bool {
		return i < 2
	}))

	shoe := Product[uint]{Id: 1, Desc: "Nike", Price: 100}
	shirt := Product[string]{Id: "_idx0012032123", Desc: "Ovejita", Price: 200}

	log.Println(shoe, shirt)
}
