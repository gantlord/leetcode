package lc1941

import "testing"

func Test_areOccurrencesEqual(t *testing.T) {
	type args struct {
		s string
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
			if got := areOccurrencesEqual(tt.args.s); got != tt.want {
				t.Errorf("areOccurrencesEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
