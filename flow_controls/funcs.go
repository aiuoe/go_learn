package flowcontrols

import "errors"

// func, struct, interface, const with capital case are exported
func Sum(a, b int) (sum int) {
	sum = a + b
	return
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division zero error")
	}

	return a / b, nil
}
