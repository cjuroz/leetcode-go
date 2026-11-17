package minimumsizesubarraysum

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{"basic", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"one", 4, []int{1, 4, 4}, 1},
		{"zero", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	solutions := map[string]func(int, []int) int{
		"MinSubArrayLen": MinSubArrayLen,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.target, tt.nums)

				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	}
}
