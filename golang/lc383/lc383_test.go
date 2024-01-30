package lc383

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "e1", args: args{ransomNote: "a", magazine: "b"}, want: false},
		{name: "e2", args: args{ransomNote: "aa", magazine: "ab"}, want: false},
		{name: "e3", args: args{ransomNote: "aa", magazine: "aab"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				//better logging
				t.Errorf("canConstruct(%v, %v) = %v, want %v", tt.args.ransomNote, tt.args.magazine, got, tt.want)
			}
		})
	}
}
