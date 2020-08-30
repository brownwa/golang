package tools

// Sqrt calculates square root using Newton's method
func Sqrt(x float64) float64 {
	z := 1.0
	for (z * z) < x {
		z -= (z*z - x) / (2 * z)
		//z -= (z*z - x/2) / (2 * z)
	}
	return z
}
