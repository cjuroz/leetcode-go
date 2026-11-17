package longestrepeatingcharacterreplacement

func CharacterReplacement(s string, k int) int {
	var seen [26]int

	left, longest, maxFreq := 0, 0, 0

	for right := 0; right < len(s); right++ {
		idx := s[right] - 'A'
		seen[idx]++

		maxFreq = max(maxFreq, seen[idx])

		if (right-left+1)-maxFreq > k {
			seen[s[left]-'A']--
			left++
		}

		longest = max(longest, right-left+1)
	}

	return longest
}
