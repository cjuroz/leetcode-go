package validpalindrome

func IsPalindrome(s string) bool {
	filtered := make([]byte, 0, len(s))

	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			filtered = append(filtered, s[i]+32)
		} else if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			filtered = append(filtered, s[i])
		}
	}

	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}

	return true
}
