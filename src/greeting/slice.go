package main

import "fmt"

func main() {
	langs := getLangs()

	for _, element := range langs {
		fmt.Println(element)
	} 
}

func getLangs() []string {
	langs := []string{"C#", "HTML", "CSS"}
	return langs
}