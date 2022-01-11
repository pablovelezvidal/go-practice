package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("The num %d is even \n", num)
			continue
		}
		fmt.Printf("The num %d is odd \n", num)
	}
}
