package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"basic", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"empty", []int{0, 1, 1}, [][]int{}},
		{"zero", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	solutions := map[string]func([]int) [][]int{
		"ThreeSum": ThreeSum,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.nums)

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
