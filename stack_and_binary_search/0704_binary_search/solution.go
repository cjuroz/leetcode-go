package binarysearch

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	mid := len(nums) / 2
	if target == nums[mid] {
		return mid
	}

	if target < nums[mid] {
		return Search(nums[:mid], target)
	}

	res := Search(nums[mid+1:], target)

	if res == -1 {
		return -1
	}

	return mid + 1 + res
}

func SearchInteractive(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1

		} else {
			right = mid - 1
		}
	}

	return -1
}
