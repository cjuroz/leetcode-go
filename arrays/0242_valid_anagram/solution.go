package validanagram

func ValidAnagramUnicode(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	smap := make(map[rune]int, len(s))

	for _, c := range s {
		smap[c]++
	}

	for _, t := range t {
		smap[t]--

		if smap[t] < 0 {
			return false
		}
	}

	return true
}
