package minimumsizesubarraysum

import (
	"math"
)

// Time: O(n), Spacke: O(1)
func MinSubArrayLen(target int, nums []int) int {
	minimalLength := math.MaxInt

	left, sum := 0, 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			minimalLength = min(minimalLength, right-left+1)

			sum -= nums[left]
			left++
		}
	}

	if minimalLength == math.MaxInt {
		minimalLength = 0
	}

	return minimalLength
}
