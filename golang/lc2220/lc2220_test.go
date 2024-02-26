package lc2220

import "testing"

func Test_minBitFlips(t *testing.T) {
	type args struct {
		start int
		goal  int
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
			if got := minBitFlips(tt.args.start, tt.args.goal); got != tt.want {
				t.Errorf("minBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
