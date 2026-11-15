package longestrepeatingcharacterreplacement

func CharacterReplacement(s string, k int) int {
	var seen [26]int

	for i := range s {
		seen[s[i]]++
	}

	return -1
}
