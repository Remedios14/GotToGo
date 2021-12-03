package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0.0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	}
	return ""
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}
	y := 1.0
	for ((y*y-x)/(2*y) > 0.00001) || ((y*y-x)/(2*y) < 0.00001) {
		y -= (y*y - x) / (2 * y)
	}
	return y, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
