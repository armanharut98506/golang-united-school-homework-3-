package main

import "fmt"

func reverse(input []int64) (result []int64) {
	for i, j := len(input) - 1, 0; i >= 0; i, j = i-1, j+1 {
		result = append(result, input[i])
	}
	return result
}

func main() {
	input := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(reverse(input))
}