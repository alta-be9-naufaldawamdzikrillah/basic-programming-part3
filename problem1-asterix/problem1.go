package main

import (
	// "bytes"
	"fmt"
)

func PlayWithAsterix(n int) string {
	// your code here
	var result string

	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			result += (" ")
		}
		for k := 0; k > i; k++ {
			result += ("*")
		}
		for l := 0; l <= i; l++ {
			result += (" *")
		}
		result += ("\n")
	}
	return result
}

func main() {
	fmt.Println("\n" + PlayWithAsterix(5))
}
