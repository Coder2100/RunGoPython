package main

import "fmt"

func main() {
	number := 0 //local variable
	increment := func(amount int) {
		number += amount
	}
	increment(1)
	increment(2)
	fmt.Println(number)
}
