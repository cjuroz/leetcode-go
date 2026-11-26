package dailytemperatures

func DailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[top] = i - top
		}
		stack = append(stack, i)
	}

	return answer
}
