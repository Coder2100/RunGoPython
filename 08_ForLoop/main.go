package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	//different way
	for i := 1; i <= 4; i++ {
		fmt.Println(i)
	}
}
