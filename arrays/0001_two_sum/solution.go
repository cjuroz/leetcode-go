package twosum

// Time: O(n²), Space: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// Time: O(n), Space: O(n)
func TwoSumHashMap(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return []int{}
}
