package twosumii

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"basic", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"basic2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"negative", []int{-1, 0}, -1, []int{1, 2}},
	}

	solutions := map[string]func([]int, int) []int{
		"TwoSum":         TwoSum,
		"TwoSumPointers": TwoSumPointers,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.numbers, tt.target)

				if slices.Compare(got, tt.want) != 0 {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	}
}
