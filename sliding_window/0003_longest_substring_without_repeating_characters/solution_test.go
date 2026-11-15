package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"basic", "abcabcbb", 3},
		{"one", "bbbbb", 1},
		{"nosubsequence", "pwwkew", 3},
	}

	solutions := map[string]func(string) int{
		"LengthOfLongestSubstring": LengthOfLongestSubstring,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.s)

				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	}
}
