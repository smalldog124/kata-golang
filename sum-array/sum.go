package sum

func Call(array []int) []int {
	length := len(array)
	var partSum []int
	for i := 0; i < length; i++ {
		var sum int
		for j := i; j < length; j++ {
			sum += array[j]
		}
		partSum = append(partSum, sum)
	}
	return append(partSum, 0)
}
