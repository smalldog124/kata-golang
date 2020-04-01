package same

func Comp(array1 []int, array2 []int) bool {
	lenthArray1 := len(array1)
	lenthArray2 := len(array2)
	if lenthArray1 == 0 || lenthArray2 == 0 {
		return false
	}
	arrayMap := makeMap(array2)
	for _, n := range array1 {
		multi := n * n
		_, ok := arrayMap[multi]
		if !ok {
			return false
		}
	}
	return true
}
func makeMap(array []int) map[int]int {
	arrayMap := make(map[int]int)
	for _, v := range array {
		arrayMap[v] = v
	}
	return arrayMap
}
