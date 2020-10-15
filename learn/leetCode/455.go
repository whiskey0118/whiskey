package main

import "fmt"

func main() {
	//a := map[int]int{}
	//b := make(map[int]int)
	//var c map[int]int
	//fmt.Println(a ==nil)
	//fmt.Println(b ==nil)
	//fmt.Println(c ==nil)

	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(a[1][1])
}

//
//func findContentChildren(g []int, s []int) int {
//	a := make(map[int]int)
//	for i:=0;i<len(s)-1;i++{
//		a[i] = g[i]
//	}
//}
