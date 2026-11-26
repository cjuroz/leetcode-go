package searchinrotatedsortedarray

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"basic-1", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"basic-2", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"basic-3", []int{1}, 0, -1},
	}

	solutions := map[string]func([]int, int) int{
		"Search": Search,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.nums, tt.target)

				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	}
}
