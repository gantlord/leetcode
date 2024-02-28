package lc2418

import (
	"reflect"
	"testing"
)

func Test_sortPeople(t *testing.T) {
	type args struct {
		names   []string
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				names:   []string{"Alice", "Bob", "Charlie", "David"},
				heights: []int{5, 6, 7, 8},
			},
			want: []string{"David", "Charlie", "Bob", "Alice"},
		},
		//["Mary","John","Emma"]
		//[5,6,7]
		{
			name: "test2",
			args: args{
				names:   []string{"Mary", "John", "Emma"},
				heights: []int{180, 165, 170},
			},
			want: []string{"Mary", "Emma", "John"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPeople(tt.args.names, tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
