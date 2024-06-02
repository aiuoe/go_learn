package _errors

import (
	"log"
)

func divider(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[danger]", r)
		}
	}()

	return a / b
}

func Panic() {
	printer := log.Println

	printer(divider(12, 2))
	printer(divider(15, 3))
	printer(divider(41, 0))
	printer(divider(22, 1))
}
