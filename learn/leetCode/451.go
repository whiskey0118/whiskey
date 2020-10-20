package main

import "fmt"

//1.判断是否包含字符串

func main() {
	a := map[int]string{1: "one", 2: "two"}
	//var a map[int]string
	if c, ok := a[1]; ok {
		fmt.Println("yes", c)
	} else {
		fmt.Println("no")
	}
}

func frequencySort(s string) string {
	tmp := make(map[rune]int)
	r := []rune(s)
	res := make([]int, len(tmp))
	for _, i := range r {
		v, ok := tmp[i]
		if ok {
			tmp[i] = v + 1
		}
	}

	for k, v := range tmp {
		res := append(res, v)
	}
	//m := make(map[rune]int)
	//ret := make([]rune,0)//存放每一个字符
	//res := make([]rune,0)//存放结果
	////注意！！要将s转成rune  因为使用for i,v:=range s的时候，这个v是rune的  最后return再转回来即可
	//r := []rune(s)
	////这个for循环是得到string中每一个字符的频率，用一个数组存下字符
	//for _,v := range r {
	//	//判断m这个map里面有没有v这个Key
	//	i,ok := m[v]
	//	if ok {//如果有的话则将m[v]的value＋1
	//		m[v] = i+1 // 即m[v]++
	//	}else{
	//		m[v] = 1
	//		ret = append(ret, v)
	//	}
	//}

}
