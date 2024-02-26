package lc2810

import "testing"

func Test_finalString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				s: "abc",
			},
			want: "abc",
		},
		{
			name: "test 2",
			args: args{
				s: "poiinter",
			},
			want: "ponter",
		},
		//"viwif"
		{
			name: "test 3",
			args: args{
				s: "viwif",
			},
			want: "wvf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalString(tt.args.s); got != tt.want {
				t.Errorf("finalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
