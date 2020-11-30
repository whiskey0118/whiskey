package main

import "fmt"

func main() {
	nums := []int{9, 3, 6, 1, 8}
	jump(nums)
}

func jump(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
}
