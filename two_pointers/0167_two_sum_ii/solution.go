package twosumii

// Time: O(n), Space: O(n)
func TwoSum(numbers []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range numbers {
		complement := target - num

		if j, ok := seen[complement]; ok {
			return []int{j + 1, i + 1}
		}
		seen[num] = i
	}

	return nil
}

// Time: O(n), Space: O(1)
func TwoSumPointers(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
