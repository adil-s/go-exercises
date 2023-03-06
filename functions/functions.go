package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func main() {
	marks := []float64{50, 50, 50, 50, 50, 60, 70, 80, 90}
	fmt.Println("average:  ", average(marks))
}
