package 27-remove-element

import (
    "sort"
    "testing"
)

func TestRemoveElement(t *testing.T) {
    testCases := []struct {
        nums        []int
        val         int
        expectedNums []int
    }{
        {
            nums:         []int{3, 2, 2, 3},
            val:          3,
            expectedNums: []int{2, 2},
        },
        {
            nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
            val:          2,
            expectedNums: []int{0, 1, 3, 0, 4},
        },
        {
            nums:         []int{1, 1, 1},
            val:          1,
            expectedNums: []int{},
        },
        {
            nums:         []int{1, 2, 3, 4, 5},
            val:          6,
            expectedNums: []int{1, 2, 3, 4, 5},
        },
    }

    for _, tc := range testCases {
        k := removeElement(tc.nums, tc.val) // Calls your implementation

        if k != len(tc.expectedNums) {
            t.Errorf("Expected length %d, got %d", len(tc.expectedNums), k)
        }

        sort.Ints(tc.nums[:k]) // Sort the first k elements of nums

        for i := 0; i < k; i++ {
            if tc.nums[i] != tc.expectedNums[i] {
                t.Errorf("At index %d, expected %d, got %d", i, tc.expectedNums[i], tc.nums[i])
            }
        }
    }
}
