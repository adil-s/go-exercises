package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		if i < 5 {
			fmt.Println("adding element using indexing", i+1)
			s[i] = i + 1
			fmt.Println("slice length: ", len(s))
			fmt.Println("slice capacity: ", cap(s))
		} else {
			fmt.Println("adding element using append ", i+1)
			s = append(s, i+1)
			fmt.Println("slice length: ", len(s))
			fmt.Println("slice capacity: ", cap(s))
		}
	}
	fmt.Println("slices length: ", len(s))
	fmt.Println("slices: ", s)
}
