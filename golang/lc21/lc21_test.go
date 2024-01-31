package lc21

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"nil", args{nil, nil}, nil},
		{"1", args{&ListNode{1, nil}, nil}, &ListNode{1, nil}},
		{"2", args{nil, &ListNode{1, nil}}, &ListNode{1, nil}},
		{"3", args{&ListNode{1, nil}, &ListNode{1, nil}}, &ListNode{1, &ListNode{1, nil}}},
		{"4", args{&ListNode{1, nil}, &ListNode{2, nil}}, &ListNode{1, &ListNode{2, nil}}},
		{"5", args{&ListNode{2, nil}, &ListNode{1, nil}}, &ListNode{1, &ListNode{2, nil}}},
		{"6", args{&ListNode{1, &ListNode{2, nil}}, &ListNode{1, nil}}, &ListNode{1, &ListNode{1, &ListNode{2, nil}}}},
		{"7", args{&ListNode{1, nil}, &ListNode{1, &ListNode{2, nil}}}, &ListNode{1, &ListNode{1, &ListNode{2, nil}}}},
		{"8", args{&ListNode{1, &ListNode{3, nil}}, &ListNode{2, nil}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
		{"9", args{&ListNode{1, nil}, &ListNode{2, &ListNode{3, nil}}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
