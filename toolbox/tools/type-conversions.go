package tools

import (
	"fmt"
	"math"
)

// ConvertTypes does some explicit type conversion
func ConvertTypes() {
	// Simplifying variable declarations
	//var x, y int = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(f)
	x := 3
	y := 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)

	fmt.Println(x, y, z)
}
