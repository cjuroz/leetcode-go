package containerwithmorewater

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"basic", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"tiny", []int{1, 1}, 1},
	}

	solutions := map[string]func([]int) int{
		"MaxArea": MaxArea,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.height)

				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	}
}
