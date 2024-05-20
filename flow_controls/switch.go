package flowcontrols

import (
	"fmt"
	"runtime"
)

func Switch() {
	// switch os := runtime.GOOS; os {
	switch os := runtime.GOOS; {
	case os == "windows":
		fmt.Println("windows:", os)
	case os == "linux":
		fmt.Println("linux:", os)
	case os == "darwin":
		fmt.Println("macOs:", os)
	default:
		fmt.Println("OTHER OS:", os)
	}
}
