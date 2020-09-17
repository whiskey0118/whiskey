//8.7
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type UserRd struct {
		Username        string
		LoginErrorTimes int
		Status          string
	}

	ud := UserRd{
		Username:        "test",
		LoginErrorTimes: 0,
		Status:          "lock",
	}
	s, _ := json.Marshal(ud)
	fmt.Printf("%s", string(s))
	var p UserRd
	json.Unmarshal([]byte(string(s)), p)
	fmt.Println(p)

}
