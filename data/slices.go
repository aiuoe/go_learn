package data

import "fmt"

func Slices() {
	week := []string{"sunday", "monday", "tuesday", "wednesday", "thursday"}

	fmt.Println("Length:", len(week))
	fmt.Println("Cap:", cap(week))

	// use append for overflow len of slice
	week = append(week, "friday")
	week = append(week, "saturday")

	fmt.Println("Length:", len(week))
	fmt.Println("Cap:", cap(week))

	weekdays := week[1:6]

	withOutTuesday := append(week[:2], week[3:]...)

	fmt.Println("Week day", weekdays)

	fmt.Println("withOutTuesday:", withOutTuesday)

	// can use copy method for copy slice inside a slice
}
