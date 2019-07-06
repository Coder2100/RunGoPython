package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (this Point) distance(other Point) float64 {
	return math.Sqrt(this.x*other.x + this.y*other.y)
}

// dince structs get automatically copied,
//its better to pass it as pinter
func (this *Point) distance_better(other *Point) float64 {
	return math.Sqrt(this.x*other.x + this.y*other.y)
}

//
func main() {
	p1 := Point{1, 3}
	p2 := Point{2, 4}
	fmt.Println(p1.distance(p2))
	fmt.Println(p1.distance_better(&p2))
}
