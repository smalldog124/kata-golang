package multi

func Call(array []int) []int {
	var result []int
	newArray := array
	newArray = append(newArray, array...)
	for i := 0; i < len(array); i++ {
		x := newArray[i+1] * newArray[i+2]
		result = append(result, x)
	}
	return result
}
