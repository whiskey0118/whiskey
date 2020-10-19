package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
	str := "client 1"
	buf := make([]byte, 1024)
	msgBack, err := conn.Write([]byte(str))
	if err != nil {
		log.Fatal(conn.RemoteAddr().String(), "conn write err: ", err)
		os.Exit(1)
	}
	msg, err := conn.Read(buf)
	if err != nil {
		log.Println("conn read err: ", err)
	}
	fmt.Println(conn.RemoteAddr().String(), "服务器反馈： ", string(buf[:msg]), msgBack, "；实际发送了", len(str))
	conn.Write([]byte("ok"))
}
