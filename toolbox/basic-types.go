package main

import (
	"fmt"
	"math/cmplx"
)

var (
	// ToBe : Represents a boolean regarding existential questions
	ToBe bool = false
	// MaxInt : Maximum value of a 64-bit unsigned int
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
