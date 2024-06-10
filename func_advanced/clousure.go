package func_advanced

import "log"

func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Clousure() {
	a := incrementer()

	log.Println(a())
	log.Println(a())
	log.Println(a())
	log.Println(a())

	b := incrementer()
	log.Println(b())
	log.Println(b())
	log.Println(b())
	log.Println(b())
}
