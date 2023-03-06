package main

import "fmt"

func swap(n1 *int, n2 *int) {
	tmp := *n1
	*n1 = *n2
	*n2 = tmp
}

func main() {
	x, y := 4, 5
	fmt.Println("swap: ", x, y)
	swap(&x, &y)
	fmt.Println("after swapping: ", x, y)
}
