package basics

import (
	"fmt"
	"strconv"
)

func Parse() {
	var type_int16 int16 = 50
	var type_int32 int32 = 100

	str := "100"
	number := 100

	i, _ := strconv.Atoi(str)
	s := strconv.Itoa(number)

	fmt.Println("Parse sum:", int32(type_int16)+type_int32)
	fmt.Printf("ParsedInt: %T \n", i)
	fmt.Printf("ParsedStr: %T \n", s)
}
