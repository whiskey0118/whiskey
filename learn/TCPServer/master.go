package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"net"
	"os"
	"time"
)

func main() {
	netListener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Error(err)
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
		ErrorNotice(err)
		//输出信息
		fmt.Println(conn.RemoteAddr().String(), " receive data: ", string(buf[:msg]))
		bufReturn := "我收到了"
		//返回信息给客户端
		msgR, err := conn.Write([]byte(bufReturn))
		ErrorNotice(err)
		fmt.Println(conn.RemoteAddr().String(), "客户端收到响应", string(buf[:msg]), "客户端收到了：", msgR)
	}
	defer conn.Close()
}

func ErrorFatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ErrorNotice(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogOut(ip string, msgRec int, msg string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	res := "time:" + t + "  |ipaddr:" + ip + "  |receive byte:" + string(msgRec) + "  |msg:" + msg
	fmt.Println(res)
}
