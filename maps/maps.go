package main

import "fmt"

func main() {
	//declaring map with fruit name as key and quantity as value
	fruits := make(map[string]int)
	fruits["apple"] = 10
	fmt.Println("map length: ", len(fruits))
	fmt.Println("fruits: ", fruits)
	detailFruits := map[string]map[string]string{
		"apple": {
			"color":    "red",
			"quantity": "10",
		},
		"banana": {
			"color":    "yellow",
			"quantity": "12",
		},
	}

	fmt.Println("detail of fruits: ", detailFruits)
	fmt.Println("apple detail: ", detailFruits["apple"])
	fmt.Println("apple color: ", detailFruits["apple"]["color"])
}
