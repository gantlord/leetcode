package lc1979

import "testing"

func Test_findGCD(t *testing.T) {
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
			if got := findGCD(tt.args.nums); got != tt.want {
				t.Errorf("findGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
