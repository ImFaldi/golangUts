package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for i, num := range nums {
		total += num
		i++
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
