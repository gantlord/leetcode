package lc530

import "testing"

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{&TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}}, 1},
		{"case2", args{&TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 7}}}, 1},
		{"case3", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 3}}}, 2},
		{"case4", args{&TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 48, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 49}}}}, 1},
		{"case5", args{&TreeNode{Val: 236, Left: &TreeNode{Val: 104, Left: &TreeNode{Val: 227, Left: &TreeNode{Val: 221}, Right: &TreeNode{Val: 235}}, Right: &TreeNode{Val: 337, Left: &TreeNode{Val: 297}, Right: &TreeNode{Val: 349}}}, Right: &TreeNode{Val: 701, Left: &TreeNode{Val: 666, Left: &TreeNode{Val: 655}, Right: &TreeNode{Val: 668}}, Right: &TreeNode{Val: 791, Left: &TreeNode{Val: 739}, Right: &TreeNode{Val: 801}}}}}, 1},
		{"case6", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}}}}, 2},
		{"case7", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}}}}, 1},
		{"case8", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}}}}, 0},
		{"case9", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}}}}, 0},
		{"case10", args{&TreeNode{Val: 543, Left: &TreeNode{Val: 384, Right: &TreeNode{Val: 445}}, Right: &TreeNode{Val: 652, Right: &TreeNode{Val: 699}}}}, 47},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
