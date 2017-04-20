package main

import (
	"math"
	"fmt"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func add(x, y float64) float64 {
	return x + y
}


func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	myAdd := add(3,4)
	fmt.Println(myAdd)
	fmt.Println(compute(hypot))
}
