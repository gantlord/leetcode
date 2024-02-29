package lc2733

import "testing"

func Test_findNonMinOrMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNonMinOrMax(tt.args.nums); got != tt.want {
				t.Errorf("findNonMinOrMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
