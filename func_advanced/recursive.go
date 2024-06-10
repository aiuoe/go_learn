package func_advanced

import "log"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Recursive() {
	result := factorial(5)

	log.Println(result)
}
