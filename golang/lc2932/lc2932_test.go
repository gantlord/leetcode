package lc2932

import "testing"

func Test_maximumStrongPairXor(t *testing.T) {
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
			if got := maximumStrongPairXor(tt.args.nums); got != tt.want {
				t.Errorf("maximumStrongPairXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
