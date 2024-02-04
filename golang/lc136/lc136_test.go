package lc136

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{2, 2, 1}}, 1},
		{"case2", args{[]int{4, 1, 2, 1, 2}}, 4},
		{"case3", args{[]int{1}}, 1},
		{"case4", args{[]int{1, 1, 2, 2, 3}}, 3},
		{"case5", args{[]int{1, 1, 2, 2, 3, 3, 4}}, 4},
		//[1,3,1,-1,3]
		{"case6", args{[]int{1, 3, 1, -1, 3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
