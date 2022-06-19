package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, len(input))
	i := 0
	for key, _ := range input {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, element := range keys {
		fmt.Println(element)
		result = append(result, input[element])
		i++
	}
	return result
}