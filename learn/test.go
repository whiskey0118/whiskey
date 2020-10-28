//8.7
package main

import (
	"fmt"
	"net/url"
)

func main() {
	a := "http://www.baidu.com/"
	res, _ := url.Parse(a)
	fmt.Println(res.Scheme)
}
