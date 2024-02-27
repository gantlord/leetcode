package lc2913

import "testing"

func Test_sumCounts(t *testing.T) {
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
			if got := sumCounts(tt.args.nums); got != tt.want {
				t.Errorf("sumCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
