package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	server := "127.0.0.1:9090"
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		log.Fatal("resolve addr err: ", err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("dial err:", err)
	}

	log.Println(conn.RemoteAddr().String(), "Connection successful")
	sender(conn)

}

func sender(conn *net.TCPConn) {
	str := os.Args
	remoteAddr := conn.RemoteAddr().String()
	buf := make([]byte, 1024)
	msgBack, err := conn.Write([]byte(str))
	ErrorNotice1(err)
	LogOut1(remoteAddr, msgBack, string(buf))
	msgBack1, err := conn.Read(buf)
	LogOut1(remoteAddr, msgBack1, string(buf))
	ErrorNotice1(err)

	conn.Write([]byte("ok"))
}

func ErrorFatal1(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ErrorNotice1(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogOut1(ip string, msgRec int, msg string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	res := "time:" + t + "  |ipaddr:" + ip + "  |receive byte:" + string(msgRec) + "  |msg:" + msg
	fmt.Println(res)
}
