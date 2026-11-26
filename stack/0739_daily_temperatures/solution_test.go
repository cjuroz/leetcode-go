package dailytemperatures

import (
	"slices"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{"basic-1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"basic-2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"basic-3", []int{30, 60, 90}, []int{1, 1, 0}},
	}

	solutions := map[string]func([]int) []int{
		"DailyTemperatures": DailyTemperatures,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.temperatures)

				if !slices.Equal(got, tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	}
}
