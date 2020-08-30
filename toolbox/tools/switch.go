package tools

import (
	"fmt"
	"runtime"
)

// SwitchThroughCases demonstrates the use of a switch statement
func SwitchThroughCases() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
