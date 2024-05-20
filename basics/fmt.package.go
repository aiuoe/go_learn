package basics

import "fmt"

func Fmt() {
	var name string
	var domain string

	/*
		verbs:
		%s: string
		%d: int
		%g: float32
		%t: boolean
		%T: type
		%+v: print struct with keys
		%p: chan pointer
	*/

	fmt.Print("Write your name: ")
	fmt.Scan(&name)

	fmt.Print("Write your domain: ")
	fmt.Scan(&domain)

	email := fmt.Sprintf("%s@%s", name, domain)

	fmt.Printf("Name: %s, Domain: %s \n", name, domain)
	fmt.Println("Email formated:", email)
}
