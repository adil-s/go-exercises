package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer file.Close()
	cnt := make([]byte, 100)
	count, err := file.Read(cnt)
	fmt.Printf("%d %s", count, cnt)
}
