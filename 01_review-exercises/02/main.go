package main

import "fmt"

type person struct {
	fname, lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println("Good Morning")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("Shaken, not stirred")
}

func main() {
	p1 := person{"Blake", "Sager"}
	fmt.Println(p1.fname)
	p1.pSpeak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(sa1.fname)
	sa1.saSpeak()
	sa1.person.pSpeak()

}
