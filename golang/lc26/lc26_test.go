package lc26

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name   string
		args   args
		k      int
		result args
	}{
		{"empty", args{[]int{}}, 0, args{[]int{}}},
		{"single", args{[]int{1}}, 1, args{[]int{1}}},
		{"two", args{[]int{1, 2}}, 2, args{[]int{1, 2}}},
		{"two dup", args{[]int{1, 1}}, 1, args{[]int{1}}},
		{"three", args{[]int{1, 2, 3}}, 3, args{[]int{1, 2, 3}}},
		{"three dup", args{[]int{1, 1, 1}}, 1, args{[]int{1}}},
		{"four", args{[]int{1, 2, 3, 4}}, 4, args{[]int{1, 2, 3, 4}}},
		{"four dup", args{[]int{1, 2, 2, 3}}, 3, args{[]int{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := removeDuplicates(tt.args.nums); got != tt.k {
				t.Errorf("removeDuplicates() = %v, want %v, result = %v", got, tt.k, tt.args.nums)
			}
			for i := 0; i < tt.k; i++ {
				if tt.args.nums[i] != tt.result.nums[i] {
					t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums, tt.result.nums)
				}
			}
		})
	}
}
