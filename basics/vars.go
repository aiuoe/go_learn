package basics

import (
	"fmt"
)

func Vars() {
	// declare var and default value of string is ""
	/*
		var (
			name, lastname string
			age int
		)
	*/

	// var name, lastname, age = "rub3n", "cortez", 31
	// name, lastname, age := "rub3n", "cortez", 31

	// declare and assign value
	name := "rub3n"

	fmt.Printf("Hello %s \n", name)
}
