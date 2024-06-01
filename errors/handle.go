package _errors

import (
	"errors"
	"log"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("error: can't divide by 0")
	}

	return a / b, nil
}

func Handler() {
	// var str string

	// fmt.Print("write a number: ")
	// fmt.Scanln(&str)

	// number, err := strconv.Atoi(str)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Number is: ", number)

	result, err := divide(12.3, 0)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)
}
