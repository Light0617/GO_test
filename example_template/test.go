package main

import "fmt"

type secretAgent struct {
	person
	licene bool
}

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, `say`, "good morning")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, `say`, "good afternoon, agent")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	xi := []int{2, 4, 5, 7}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Penny",
	}

	fmt.Println(p1)

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	sa1.speak()

	saySomething(p1)
	saySomething(sa1)
}
