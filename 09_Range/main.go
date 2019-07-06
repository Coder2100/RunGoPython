package main

import "fmt"

func main() {
	nameLists := []string{
		"Zukile",
		"Nqwesh",
		"Odwa",
		"Sam",
		"Aphelele",
	}
	for i, name := range nameLists {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}
