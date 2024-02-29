package lc2500

import "testing"

func Test_deleteGreatestValue(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 0",
			args: args{
				grid: [][]int{
					{1, 2, 4},
					{3, 3, 1},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteGreatestValue(tt.args.grid); got != tt.want {
				t.Errorf("deleteGreatestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
