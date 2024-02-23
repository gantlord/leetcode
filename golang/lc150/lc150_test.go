package lc150

import "testing"

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]string{"2", "1", "+", "3", "*"}}, 9},
		{"test2", args{[]string{"4", "13", "5", "/", "+"}}, 6},
		{"test3", args{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, 22},
		{"test4", args{[]string{"18"}}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}