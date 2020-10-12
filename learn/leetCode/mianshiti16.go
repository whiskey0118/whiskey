package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 6, 5, 11, 16, 8, 4, 10}
	var a, b int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			a = i
			fmt.Println(nums[i])
			break
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[len(nums)-i-2] > nums[len(nums)-i-1] {
			b = len(nums) - i - 1
			break
		}
	}
	fmt.Printf("%d and %d", a, b)
}

//func subSort(array []int) []int {
//
//}
