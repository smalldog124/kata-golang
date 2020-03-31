package same

func Comp(array1 []int, array2 []int) bool {
	lenth := len(array1) - 1
	for i := 0; i < lenth; i++ {
		n := array1[i]
		multi := n * n
		if multi != array2[i+1] {
			return false
		}
	}
	return true
}
