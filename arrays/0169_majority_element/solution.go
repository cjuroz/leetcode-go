package majorityelement

// Time: O(n), Space: O(n)
func MajorityElementMap(nums []int) int {
	msize := len(nums) / 2

	counts := make(map[int]int, msize)

	for _, num := range nums {
		counts[num]++
		if counts[num] > msize {
			return num
		}
	}

	return -1
}

// Time: O(n), Space: O(1)
func MajorityElement(nums []int) int {
	candidate, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}

	return candidate
}
