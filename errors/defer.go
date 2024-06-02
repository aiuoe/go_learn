package _errors

import (
	"fmt"
	"os"
)

func Defer() {
	file, err := os.Create("data.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hello World"))

	if err != nil {
		fmt.Println(err)
		return
	}
}
