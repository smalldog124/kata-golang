package same

func Comp(array1 []int, array2 []int) bool {
	if len(array1) == 0 || len(array2) == 0 {
		return false
	}
	lenth := len(array1) - 1
	for i := 0; i < lenth; i++ {
		n := array1[i]
		multi := n * n
		if multi != array2[i+1] {
			return false
		}
	}
	return (array1[lenth] * array1[lenth]) == array2[0]
}
