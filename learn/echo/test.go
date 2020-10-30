package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

type Arg []string

func (c *Arg) String() string {
	return strings.Join([]string(*c), " ")
}

// Set is the method flag package calls
func (c *Arg) Set(value string) error {
	*c = append(*c, value)
	return nil
}

func main() {

	//haha := flag.String("haha","rest","sdfsdf")
	var conf Arg
	flag.Var(&conf, "haha", "test")
	flag.Parse()
	log.SetFlags(0)
	fmt.Println(conf)
}

func test(nums []int) int {
	a := 0
	for i := range nums {
		a ^= i
	}
	return a
}
