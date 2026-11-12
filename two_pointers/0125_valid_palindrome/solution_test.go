package validpalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"basic", "A man, a plan, a canal: Panama", true},
		{"nope", "race a car", false},
		{"empty", " ", true},
	}

	solutions := map[string]func(string) bool{
		"IsPalindrome": IsPalindrome,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.s)

				if got != tt.want {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	}
}
