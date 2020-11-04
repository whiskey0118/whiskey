package main

func main() {

}

//双指针

func trap(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if leftMax < height[left] {
				leftMax = height[left]
			} else {
				res = res + leftMax - height[left]
			}
			left++
		} else {
			if rightMax <= height[right] {
				rightMax = height[right]
			} else {
				res = res + rightMax - height[right]
			}
			right--
		}
	}
	return res
}
