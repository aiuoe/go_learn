package data

import (
	"fmt"
	"strings"
)

type String struct {
	Value string
}

func (s *String) ToUpper() string {
	return strings.ToUpper(s.Value)
}

func Pointers() {
	// vars pass as value
	a := 10

	// is pass by ref b change value when a change value
	// for pass by ref use & before var name
	var b *int = &a

	a = 20

	// use * for unreference
	fmt.Println(*b)

	str := String{Value: "Hello World!"}

	fmt.Println("To Upper:", str.ToUpper())
}
