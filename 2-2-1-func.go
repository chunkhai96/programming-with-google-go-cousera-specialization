package main

import (
	"fmt"
)

func main() {

	var a, v0, s0, t float64

	fmt.Printf("Enter acceleration (a): ")
	fmt.Scan(&a)
	fmt.Printf("Enter initial velocity (v0): ")
	fmt.Scan(&v0)
	fmt.Printf("Enter initial displacement (s0): ")
	fmt.Scan(&s0)
	fn := GenDisplaceFn(a, v0, s0)
	for {
		fmt.Printf("Enter time (t): ")
		fmt.Scan(&t)
		fmt.Printf("Displacement (s) = %.2f\n", fn(t))
	}
}

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	return func (t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}
}