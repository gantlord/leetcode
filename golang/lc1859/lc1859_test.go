package lc1859

import "testing"

func Test_sortSentence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSentence(tt.args.s); got != tt.want {
				t.Errorf("sortSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
