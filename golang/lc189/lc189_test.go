package lc189

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		array    args
		k        int
		expected args
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 3, args{[]int{5, 6, 7, 1, 2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.array.nums, tt.k)
		})
		for i := 0; i < len(tt.array.nums); i++ {
			if tt.array.nums[i] != tt.expected.nums[i] {
				t.Errorf("At index %d, expected %d, got %d", i, tt.expected.nums[i], tt.array.nums[i])
			}
		}
	}
}
