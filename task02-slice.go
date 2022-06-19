package homework

func reverse(input []int64) (result []int64) {
	result := make([]int, len(input))
	for i, j := len(input) - 1, 0; i >= 0; i, j = i-1, j+1 {
		result[j] = input[i]
	}
	return result
}
