package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"log"
	"net"
)

func main() {
	netListener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		logs.Error(err)
	}
	defer netListener.Close()

	for {
		conn, err := netListener.Accept()
		if err != nil {
			log.Fatal("conn err", err)
		}
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		msg, err := conn.Read(buf)
		if err != nil {
			log.Fatal("conn read err: ", err)
			return
		}
		//输出信息
		fmt.Println(conn.RemoteAddr().String(), " receive data: ", string(buf[:msg]))
		bufReturn := "我收到了"
		//返回信息给客户端
		msgR, err := conn.Write([]byte(bufReturn))
		if err != nil {
			log.Println(conn.RemoteAddr().String(), " client 没有收到响应")
		}
		fmt.Println(conn.RemoteAddr().String(), "客户端收到响应", string(buf[:msg]), "客户端收到了：", msgR)

	}
}
