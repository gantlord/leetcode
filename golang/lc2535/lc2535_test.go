package lc2535

import "testing"

func Test_differenceOfSum(t *testing.T) {
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
			if got := differenceOfSum(tt.args.nums); got != tt.want {
				t.Errorf("differenceOfSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
