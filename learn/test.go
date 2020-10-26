//8.7
package main

import (
	"fmt"
	"os"
)

func main() {
	c, err := os.Open("F:\\我的\\go\\whiskey\\learn\\test1.go")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[:n]))
		fmt.Println("_____")
	}

}
