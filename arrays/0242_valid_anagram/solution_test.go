package validanagram

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"true", "anagram", "nagaram", true},
		{"false", "rat", "car", false},
		{"length", "hi", "ihh", false},
	}

	solutions := map[string]func(string, string) bool{
		"ValidAnagram": ValidAnagramUnicode,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.s, tt.t)

				if got != tt.want {
					t.Errorf("got %t, want %t", got, tt.want)
				}
			})
		}
	}
}
