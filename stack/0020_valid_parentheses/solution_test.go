package validparentheses

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"basic-1", "()", true},
		{"basic-2", "()[]{}", true},
		{"false", "(]", false},
		{"basic-3", "([])", true},
		{"false-2", "([)]", false},
	}

	solutions := map[string]func(string) bool{
		"IsValid":    IsValid,
		"IsValidMap": IsValidMap,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.s)

				if got != tt.want {
					t.Errorf("got %t, want %t", got, tt.want)
				}
			})
		}
	}
}
