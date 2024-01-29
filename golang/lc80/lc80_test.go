package lc80

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		wantk    int
		wantargs args
	}{
		{"1", args{[]int{1, 1, 1, 2, 2, 3}}, 5, args{[]int{1, 1, 2, 2, 3}}},
		{"2", args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7, args{[]int{0, 0, 1, 1, 2, 3, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.wantk {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.wantk)
			}
			for i := 0; i < tt.wantk; i++ {
				if tt.args.nums[i] != tt.wantargs.nums[i] {
					t.Errorf("At index %d, expected %d, got %d", i, tt.wantargs.nums[i], tt.args.nums[i])
				}
			}
		})
	}
}
