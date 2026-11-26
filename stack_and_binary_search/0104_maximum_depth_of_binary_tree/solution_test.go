package maximumdepthofbinarytree

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"basic-1", &TreeNode{3,
			&TreeNode{9, nil, nil},
			&TreeNode{20,
				&TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil}}}, 3},
		{"basic-2", &TreeNode{1, nil,
			&TreeNode{2, nil, nil}}, 2},
	}

	solutions := map[string]func(*TreeNode) int{
		"MaxDepth": MaxDepth,
	}

	for _, tt := range tests {
		for name, solve := range solutions {
			t.Run(tt.name+"/"+name, func(t *testing.T) {
				got := solve(tt.root)

				if got != tt.want {
					t.Errorf("got %d, want %d", got, tt.want)
				}
			})
		}
	}
}
