package data

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  string
}

func Structs() {
	// You can also start a Struct with New
	cat := Animal{"copito", "5 meses"}

	fmt.Println("Cat:", cat)
}
