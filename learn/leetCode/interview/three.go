package main

import "fmt"

//给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
//func countSmaller(nums []int) []int {
//	for i:=0;i<len(nums)-1;i++{
//
//	}
//}

func main() {
	nums := []int{2, 15, 4, 10, 6, 51, 1}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				//nums[i] = nums[j]^nums[i]
				//nums[j] = nums[j]^nums[i]
				//nums[i] = nums[j]^nums[i]
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
}
