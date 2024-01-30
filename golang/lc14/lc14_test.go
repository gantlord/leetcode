package lc14

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//		{"empty", args{[]string{}}, ""},
		{"one-one", args{[]string{"a"}}, "a"},
		{"three-two", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"short", args{[]string{"ab", "a"}}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
