package main

import "fmt"

func main() {
	gopher1Name := "Rohan"
	gopher1Age := 21

	gopher2Age := 90
	gopher2Name := "Noodles"

	if gopher1Age < 60 {
		highjump(gopher1Name)
	} else {
		lowjump(gopher2Name)
	}

	if gopher2Age < 60 {
		highjump(gopher2Name)
	} else {
		lowjump(gopher2Name)
	}
}

func highjump(name string) {
	fmt.Println(name, "can jump pretty high")
}

func lowjump(name string) {
	fmt.Println(name, "can still jump though")
}