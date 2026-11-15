package besttimetobuyandsellstock

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"basic", []int{7, 1, 5, 3, 6, 4}, 5},
		{"nope", []int{7, 6, 4, 3, 1}, 0},
	}

	solutions := map[string]func([]int) int{
		"MaxProfitBruteForce": MaxProfitBruteForce,
		"MaxProfit":           MaxProfit,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.prices)

				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	}
}
