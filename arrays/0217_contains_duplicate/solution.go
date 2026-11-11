package containsduplicate

func ContainsDuplicateHashMap(nums []int) bool {
	duplicates := make(map[int]bool)

	for _, num := range nums {
		if duplicates[num] {
			return true
		}
		duplicates[num] = true
	}

	return false
}

func ContainsDuplicateHashSet(nums []int) bool {
	duplicates := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := duplicates[num]; exists {
			return true
		}
		duplicates[num] = struct{}{}
	}

	return false
}
