package main

import "fmt"

func main() {
	/*var langs [3]string
	langs[0] = "Go"
	langs[1] = "Java"
	langs[2] = "Java Script"
	//langs[3] = "Lol"
	fmt.Println(langs)
	fmt.Println("\nSlice:\n")*/
	printSlice()
}

func printSlice() {
	//var langs[]string
	/*langs = append(langs, "Lol")
	langs = append(langs, "Ruby")
	langs = append(langs, "Python")
	langs = append(langs, "Kotlin")*/
	langs := []string{"Lol", "Python", "Kotlin", "Ruby"}
	fmt.Println(langs)
}