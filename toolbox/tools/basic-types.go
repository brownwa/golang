// Package tools contains coding examples for practical usage of Go
package tools

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

// DescribeTypes describes some type names and example values
func DescribeTypes() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
