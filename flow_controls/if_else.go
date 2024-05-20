package flowcontrols

import (
	"fmt"
	"time"
)

func IfElse() {
	// now := time.Now()
	// hour := now.Hour()

	if now := time.Now(); now.Hour() < 12 {
		fmt.Println("morning")
	} else if now.Hour() < 17 {
		fmt.Println("afternoon")
	} else {
		fmt.Println("night")
	}

	fmt.Println("Hour:", time.Now().Hour())
}
