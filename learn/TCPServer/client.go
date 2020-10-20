package main

import (
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
	str := "This is client 1"
	remoteAddr := conn.RemoteAddr().String()
	buf := make([]byte, 1024)
	msgBack, err := conn.Write([]byte(str))
	CheckError(err)
	LogOut(remoteAddr, msgBack, string(buf))
	msgBack1, err := conn.Read(buf)
	LogOut(remoteAddr, msgBack1, string(buf))
	CheckError(err)

	conn.Write([]byte("ok"))
}
