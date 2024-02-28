package lc2778

import "testing"

func Test_sumOfSquares(t *testing.T) {
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
			if got := sumOfSquares(tt.args.nums); got != tt.want {
				t.Errorf("sumOfSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
