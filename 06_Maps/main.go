package main

import "fmt"

func main() {
	elements := make(map[string]int)
	elements["H"] = 1
	fmt.Println(elements["H"])
	//Remove by key
	elements["0"] = 8
	//	delete(elements, "0")
	//Only do something with the element if its in the map
	if number, ok := elements["0"]; ok {
		fmt.Println(number) //
	}

	if number, ok := elements["H"]; ok {
		fmt.Println(number) //1
	}
}
