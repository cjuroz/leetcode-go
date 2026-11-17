package longestrepeatingcharacterreplacement

import (
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"basic", "ABAB", 2, 4},
		{"baisc2", "AABABBA", 1, 4},
	}

	solutions := map[string]func(string, int) int{
		"CharacterReplacement": CharacterReplacement,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.s, tt.k)

				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	}
}
