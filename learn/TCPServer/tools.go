package main

import (
	"fmt"
	"os"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func LogOut(ip string, msgRec int, msg string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	res := "time:" + t + "  |ipaddr:" + ip + "  |receive byte:" + string(msgRec) + "  |msg:" + msg
	fmt.Println(res)
}
