//8.7
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	a := sha1.New()
	b, err := a.Write([]byte("sdf"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Printf("%x", a.Sum(nil))

}
