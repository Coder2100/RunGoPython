package main

import "fmt"

func main() {
	//initialize array
	var numbers [5]int
	// change the array
	numbers[2] = 100
	//getnumber
	getNumber := numbers[1:3]
	fmt.Println(getNumber)
	//length of it
	fmt.Println(len(numbers))
	//initialize a slice
	var scores []float64
	scores = append(scores, 1.1) //recreate to append
	scores[0] = 2.2              // change the values of an array
	fmt.Println(scores)
	// strng arrays
	var listofPeople [100]string
	listofPeople[0] = "Dan"
	listofPeople[1] = "Ducks"
	listofPeople[2] = "Zukile"
	fmt.Println(len(listofPeople))
	//output the list
	fmt.Println(listofPeople[0])
	fmt.Println(listofPeople[1])
	fmt.Println(listofPeople[2])

}
