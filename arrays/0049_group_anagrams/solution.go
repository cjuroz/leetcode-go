package groupanagrams

import "slices"

// Time: O(n), Space: O(n)
func GroupAnagrams(strs []string) [][]string {

	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for i := 0; i < len(s); i++ {
			count[s[i]-'a']++
		}
		groups[count] = append(groups[count], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Time: O(n), Space: O(n)
func GroupAanagramsUnicode(strs []string) [][]string {

	groups := make(map[string][]string)

	for _, s := range strs {
		runes := []rune(s)
		slices.Sort(runes)
		key := string(runes)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
