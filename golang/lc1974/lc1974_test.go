package lc1974

import "testing"

func Test_minTimeToType(t *testing.T) {
	type args struct {
		word string
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
			if got := minTimeToType(tt.args.word); got != tt.want {
				t.Errorf("minTimeToType() = %v, want %v", got, tt.want)
			}
		})
	}
}
