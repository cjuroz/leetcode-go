package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"basic", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"duplicates", []int{3, 3}, 6, []int{0, 1}},
		{"negative", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
	}

	solutions := map[string]func([]int, int) []int{
		"BruteForce": TwoSumBruteForce,
		"HashMap":    TwoSumHashMap,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.nums, tt.target)

				if !slices.Equal(got, tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	}
}
