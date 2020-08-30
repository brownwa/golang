package tools

import (
	"fmt"
	"math"
)

func sqrt(x float64) (string, string) {
	if x < 0 {
		return sqrt(-x)
	}
	return fmt.Sprint(math.Sqrt(x)), "i"
}

// GetMultipleResults demonstrates multiple return values from a function
func GetMultipleResults() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))
}
