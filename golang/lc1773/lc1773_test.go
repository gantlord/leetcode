package lc1773

import "testing"

func Test_countMatches(t *testing.T) {
	type args struct {
		items     [][]string
		ruleKey   string
		ruleValue string
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
			if got := countMatches(tt.args.items, tt.args.ruleKey, tt.args.ruleValue); got != tt.want {
				t.Errorf("countMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
