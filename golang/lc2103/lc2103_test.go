package lc2103

import "testing"

func Test_countPoints(t *testing.T) {
	type args struct {
		rings string
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
			if got := countPoints(tt.args.rings); got != tt.want {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
