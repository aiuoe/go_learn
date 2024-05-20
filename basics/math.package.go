package basics

import (
	"fmt"
	"math"
)

func Math() {
	Pi := math.Pi

	fmt.Printf("Pi: %.2f \n", Pi)

	fmt.Println("Sqrt 64:", math.Sqrt(64))
	fmt.Println("Pow 2^3:", math.Pow(2, 3))

	var base, height, hypo float64

	fmt.Print("Write base:")
	fmt.Scanln(&base)

	fmt.Print("Write height:")
	fmt.Scan(&height)

	fmt.Println("Area of the Triangle:", (base*height)/2)

	hypo = math.Hypot(base, height)

	fmt.Printf("Hypotenusa of the Triangle: %.2f \n", hypo)
	fmt.Printf("Perimeter of the Triangle: %.3f \n", base+height+hypo)
}
