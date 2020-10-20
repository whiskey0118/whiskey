package main

import (
	"fmt"
	"sort"
)

func main() {
	a := map[string]int{"four": 4, "one": 1}
	//sort.Sort(a)
	s1 := []string{"c", "a", "b", "2", "1"}
	sort.Strings(s1)
	fmt.Println(s1)
	fmt.Println(a["one"])
	fmt.Println(len(a))

}
