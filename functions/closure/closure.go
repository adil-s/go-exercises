package main

import "fmt"

func makeOddNumGenerator() func() int {
	i := 1
	return func() int {
		i += 2
		return i
	}
}

func main() {
	nextOdd := makeOddNumGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
