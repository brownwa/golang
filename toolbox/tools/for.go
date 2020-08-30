package tools

import "fmt"

// Iterate demonstrates usage of a for loop
func Iterate() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
