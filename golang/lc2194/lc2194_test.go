package lc2194

import (
	"reflect"
	"testing"
)

func Test_cellsInRange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellsInRange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cellsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
