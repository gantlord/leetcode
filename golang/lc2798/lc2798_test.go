package lc2798

import "testing"

func Test_numberOfEmployeesWhoMetTarget(t *testing.T) {
	type args struct {
		hours  []int
		target int
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
			if got := numberOfEmployeesWhoMetTarget(tt.args.hours, tt.args.target); got != tt.want {
				t.Errorf("numberOfEmployeesWhoMetTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
