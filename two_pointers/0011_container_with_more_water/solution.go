package containerwithmorewater

func MaxArea(height []int) int {
	var maxArea int

	left, right := 0, len(height)-1

	for left < right {
		area := (right - left) * min(height[left], height[right])

		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
