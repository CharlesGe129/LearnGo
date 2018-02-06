package main

// Source: https://tour.golang.org/flowcontrol/8

import (
"fmt"
)

func Sqrt(x float64) float64 {
	// z -= (z*z - x) / (2*z)
	var z float64 = 1
	for i := 0; i < 10; i ++ {
		z -= (z * z - x) / (2 * z)
		fmt.Printf("%v: %f\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2));
}