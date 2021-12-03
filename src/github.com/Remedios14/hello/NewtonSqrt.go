package main

import (
	"fmt"
)

func Sqrt(x float64) (y float64) {
	y = 1.0
	for ((y*y-x)/(2*y) > 0.00001) || ((y*y-x)/(2*y) < 0.00001) {
		y -= (y*y - x) / (2 * y)
	}
	return
}

func main() {
	fmt.Println(Sqrt(2))
}
