package basics

import "fmt"

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const Pi float32 = 3.1415

func Consts() {
	// const: never change values
	// const: define in package level
	fmt.Println("PI:", Pi)
	fmt.Println("Friday:", Friday)
}
