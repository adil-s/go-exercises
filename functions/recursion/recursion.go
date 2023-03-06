package main

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + sum(n-1)
	}
}

func main() {
	fmt.Println(sum(5))
}
