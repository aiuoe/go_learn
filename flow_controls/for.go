package flowcontrols

import "fmt"

func For() {
	fruits := []string{"banana", "apple", "pera"}

	// index, value
	for _, fruit := range fruits {
		if fruit == "apple" {
			continue
		}

		fmt.Println("Fruit:", fruit)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("I:", i)

		if i == 5 {
			break
		}
	}
}
