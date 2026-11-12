package topkfrequentelements

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"basic", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"one", []int{1}, 1, []int{1}},
		{"long", []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
	}

	solutions := map[string]func([]int, int) []int{
		"TopKFrequent":           TopKFrequent,
		"TopKFrequentBucketSort": TopKFrequentBucketSort,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.nums, tt.k)

				sort.Ints(got)
				sort.Ints(tt.want)

				// Handle nil vs empty slice edge case
				if len(got) == 0 && len(tt.want) == 0 {
					return
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	}
}
