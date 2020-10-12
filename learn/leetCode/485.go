package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	var str string
	for i := 0; i < len(nums); i++ {
		str = str + strconv.Itoa(nums[i])
	}
	fmt.Println(str)
	fmt.Println(strings.Split(str, "0"))
	b := strings.Split(str, "0")
	c, _ := strconv.Atoi(b[1])
	fmt.Printf("%d", c)
	res, _ := strconv.Atoi(b[0])
	for i := 0; i < len(b)-1; i++ {
		two, _ := strconv.Atoi(b[i])
		if two > res {
			res = two
		}
	}

	fmt.Println(res)
}

func findMaxConsecutiveOnes(nums []int) int {
	i, j := 0, 0
	max := 0
	for j < len(nums) {
		if nums[j] != 1 {
			//i表示0的数量
			i = j + 1
		}
		j++
		if j-i > max {
			max = j - i
		}
	}
	return max
}
