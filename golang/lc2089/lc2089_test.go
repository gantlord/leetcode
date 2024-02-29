package lc2089

import (
	"reflect"
	"testing"
)

func Test_targetIndices(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := targetIndices(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("targetIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
