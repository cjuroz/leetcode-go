package groupanagrams

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{"basic",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"empty",
			[]string{""},
			[][]string{{""}}},
		{"one",
			[]string{"a"},
			[][]string{{"a"}}},
	}

	solutions := map[string]func([]string) [][]string{
		"GroupAnagrams":        GroupAnagrams,
		"GroupAnagramsUnicode": GroupAanagramsUnicode,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.strs)

				normalize(got)
				normalize(tt.want)

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

func normalize(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		return strings.Join(groups[i], ",") < strings.Join(groups[j], ",")
	})
}
