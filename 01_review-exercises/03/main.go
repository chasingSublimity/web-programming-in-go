package main

import "fmt"

type person struct {
	fname, lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println("Good Morning")
}

func (sa secretAgent) speak() {
	fmt.Println("Shaken, not stirred")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Blake", "Sager"}
	saySomething(p1)

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(sa1)

}
