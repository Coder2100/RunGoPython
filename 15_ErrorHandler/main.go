package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Error was: ", recover())
	}()
	panic("Mayhem!!!")
}
