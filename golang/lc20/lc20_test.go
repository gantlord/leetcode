package lc20

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, true},
		{"()", args{"()"}, true},
		{"()[]{}", args{"()[]{}"}, true},
		{"(]", args{"(]"}, false},
		{"([)]", args{"([)]"}, false},
		{"{[]}", args{"{[]}"}, true},
		{"[", args{"["}, false},
		{"]", args{"]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
