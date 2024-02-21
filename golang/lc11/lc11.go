package lc11

func maxArea(height []int) int {
	maxArea := 0
	l, r := 0, len(height)-1
	for {
		if l == r {
			return maxArea
		}
		h := min(height[l], height[r])
		a := h * (r - l)
		if a > maxArea {
			maxArea = a
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
}
