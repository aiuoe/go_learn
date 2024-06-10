package poo

import "fmt"

type Animal interface {
	Sound()
}

func Sounder(a Animal) {
	a.Sound()
}

type Dog struct {
	Name string
}

func (d *Dog) Sound() {
	fmt.Printf("I, %s, and sound as Woof! \n", d.Name)
}

type Cat struct {
	Name string
}

func (c *Cat) Sound() {
	fmt.Printf("I, %s, and sound as Meow! \n", c.Name)
}

func ExecInterface() {
	dog := &Dog{Name: "Coqueta"}
	cat := &Cat{Name: "Copito"}

	Sounder(dog)
	Sounder(cat)
}
