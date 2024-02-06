package lc134

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}}, 3},
		{"case2", args{[]int{5, 5, 1, 3, 4}, []int{8, 1, 7, 1, 1}}, 3},
		{"case3", args{[]int{1, 2, 3, 4, 5, 5, 70}, []int{2, 3, 4, 3, 9, 6, 2}}, 6},
		{"case4", args{[]int{2, 3, 4}, []int{3, 4, 3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
