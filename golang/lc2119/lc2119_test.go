package lc2119

import "testing"

func Test_isSameAfterReversals(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameAfterReversals(tt.args.num); got != tt.want {
				t.Errorf("isSameAfterReversals() = %v, want %v", got, tt.want)
			}
		})
	}
}
