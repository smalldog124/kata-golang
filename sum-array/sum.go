package sum

func Call(array []int) []int {
	length := len(array) - 1
	partSum := array
	var sum int
	for i := length; i >= 0; i-- {
		sum += array[i]
		partSum[i] = sum
	}

	return append(partSum, 0)
}
