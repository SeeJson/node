package algo

func towSum(arr []int, k int) []int {
	hashTable := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if p, ok := hashTable[k-arr[i]]; ok {
			return []int{i, p}
		}
		hashTable[arr[i]] = i
	}
	return nil
}
