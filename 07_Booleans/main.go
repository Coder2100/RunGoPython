package main

import "fmt"

func main() {
	x := 1
	if x != 0 {
		fmt.Println("Yes")
	}

	var y []string
	if len(y) != 0 {
		fmt.Println("This will not output")
	}
	fmt.Println(true && false) // false
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(!true)
	fmt.Println(!false)

	if len(y) != 0 {
		fmt.Println("This will not output")
	}
}

//x := 1
