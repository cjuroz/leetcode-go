package containsduplicate

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"basic", []int{1, 2, 3, 1}, true},
		{"none", []int{1, 2, 3, 4}, false},
		{"multiple", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	solutions := map[string]func([]int) bool{
		"HashMap": ContainsDuplicateHashMap,
		"HashSet": ContainsDuplicateHashSet,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.nums)

				if got != tt.want {
					t.Errorf("got %t, want %t", got, tt.want)
				}
			})
		}
	}
}
