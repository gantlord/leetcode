package lc1812

import "testing"

func Test_squareIsWhite(t *testing.T) {
	type args struct {
		coordinates string
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
			if got := squareIsWhite(tt.args.coordinates); got != tt.want {
				t.Errorf("squareIsWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}