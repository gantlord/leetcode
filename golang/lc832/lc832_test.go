package lc832

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	type args struct {
		image [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.args.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
