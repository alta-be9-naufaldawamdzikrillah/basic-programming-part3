package main

import "fmt"

func MeanMedian(arrayInput []float64) (float64, float64) {
	// your code here
	var sum float64 = 0
	for i := 0; i < len(arrayInput); i++{
		sum = sum + arrayInput[i]
	}
	var mean float64 = sum / float64(len(arrayInput))
	var median float64
	tampung := len(arrayInput) / 2
	if len(arrayInput)%2 == 0{
		median = (arrayInput[tampung - 1]) + arrayInput[tampung]/2
	}else {
		median = arrayInput[tampung]
	}
	return mean, median
}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))       // 3 3
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))     // 10.4 9
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}
