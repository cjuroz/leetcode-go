package longestsubstringwithoutrepeatingcharacters

func LengthOfLongestSubstring(s string) int {
	length, longest := 0, 0
	var seen [128]bool

	begin, cursor := 0, 0
	for cursor < len(s) {
		c := s[cursor]
		length = cursor - begin + 1
		if seen[c] {
			seen[s[begin]] = false
			begin++
		} else {
			seen[c] = true
			longest = max(longest, length)
			cursor++
		}
	}

	return longest
}
