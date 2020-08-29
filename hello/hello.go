package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

// Run the main Go code.
func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hellow World", "Hello Go"))
}
