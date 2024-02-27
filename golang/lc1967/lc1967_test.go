package lc1967

import "testing"

func Test_numOfStrings(t *testing.T) {
	type args struct {
		patterns []string
		word     string
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
			if got := numOfStrings(tt.args.patterns, tt.args.word); got != tt.want {
				t.Errorf("numOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
