package func_advanced

import "log"

func sum(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		log.Println(number)
		total += number
	}

	return total
}

func Variadic() {
	result := sum(1, 2.3, 3, 4.50)

	log.Println(result)
}
