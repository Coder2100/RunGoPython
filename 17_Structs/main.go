package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func distance(point1 Point, point2 Point) float64 {
	return math.Sqrt(point1.x*point2.x + point1.y*point2.y)
}

//since struckt get automatically copied, its better to pass reults to the pointer
func distanceBetter(point1 *Point, point2 *Point) float64 {
	return math.Sqrt(point1.x*point2.x + point1.y*point2.y)
}

func main() {
	p1 := Point{1, 3}
	p2 := Point{2, 4}
	fmt.Println(distance(p1, p2))
	fmt.Println(distanceBetter(&p1, &p2))
}
