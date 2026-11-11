package majorityelement

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"basic", []int{3, 2, 3}, 3},
		{"basic2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	solutions := map[string]func([]int) int{
		"MajorityElement":    MajorityElement,
		"MajorityElementMap": MajorityElementMap,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.nums)

				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	}
}
