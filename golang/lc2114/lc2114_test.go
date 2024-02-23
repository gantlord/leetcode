package lc2114

import "testing"

func Test_mostWordsFound(t *testing.T) {
	type args struct {
		sentences []string
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
			if got := mostWordsFound(tt.args.sentences); got != tt.want {
				t.Errorf("mostWordsFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
