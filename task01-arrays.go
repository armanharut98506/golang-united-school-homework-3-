package homework

func average(input [15]float32) (result float32) {
	for _, element := range input {
		result += element
	}
	return result / float32(len(input))
}
