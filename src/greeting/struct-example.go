package main

import "fmt"

type gopher struct {
	name string
	age int
	isAdult bool
}

func (g gopher) jump() string {
	if g.age < 65 {
		return g.name + " can jump high!!"
	} else {
		return g.name + " can still jump"
	}
}

func main() {
	gopher1 := &gopher{name: "Phil", age: 30}
	gopher2 := gopher{name: "Noodles", age: 90}
	validateAge(gopher1) 
	// when & is not used, copy of original struct is created.
	// and sent to validateAge()

	fmt.Println(gopher1)

	fmt.Println(gopher1.jump())
	fmt.Println(gopher2.jump())
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}