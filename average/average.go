package main

import "fmt"

func main() {
	var total float64 = 0
	marks := []float64{45, 50, 100, 100, 50}
	fmt.Println("number of subjects: ", len(marks))
	fmt.Println("marks: ", marks)
	for _, value := range marks {
		total += value
	}
	fmt.Println("total: ", total)
	fmt.Println("average: ", total/float64(len(marks)))
}
