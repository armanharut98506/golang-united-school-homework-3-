package homework

func average(input [15]float32) (result float32) {
	sum := 0
	for _, element := range input {
		sum += element
	}
	return sum / len(input)
}
