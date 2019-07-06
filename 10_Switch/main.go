package main

import (
	"fmt"
	"strconv"
)

//create method to accept user input
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Not a number")
	}
	return i
}

func main() {
	var numberString string
	fmt.Scanln(&numberString)
	number := stringToInt(numberString)

	switch number {
	case 1:
		fmt.Println("Sabelo")

	case 2:
		fmt.Println("Ntobeko")
	case 3:
		fmt.Println("Monwabisi")
	case 4:
		fmt.Println("Zukile")
	default:
		fmt.Printf("I have no idea what is %d", number)
	}
}
