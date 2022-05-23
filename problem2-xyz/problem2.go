package main

import "fmt"

func DrawXYZ(n int) string {
	// your code here

	var result string
	for i := 1; i <= n*n; i++ {
		if i%3 == 0 {
			result += "X"
		} else if i%2 == 0 {
			result += "Z"
		} else {
			result += "Y"
		}
		if i%n == 0 {
			result += "\n"
		}
	}
	return result
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}
