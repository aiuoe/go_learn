package data

import "fmt"

func Array() {
	// unkown size ? use ... instead number
	var arr = [...]int{1, 2, 3, 4, 5}
	arr[0] = 100
	arr[1] = 200

	for i := 0; i < len(arr); i++ {
		fmt.Printf("{index: %d, value: %v} \n", i, arr[i])
	}

	// index, value
	for index, value := range arr {
		fmt.Printf("{index: %d, value: %v}\n", index, value)
	}
}
