package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	netListener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
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
	remoteAddr := conn.RemoteAddr().String()
	defer conn.Close()
	for {
		bufNum, err := conn.Read(buf)
		//输出信息
		if err != nil {
			//如果客户端关闭，退出函数，继续接受连接，如果不退出函数，整个应用将关闭
			log.Println("client has close")
			return
		}
		LogOut(remoteAddr, 0, string(buf[:bufNum]))
		bufReturn := "master had receive"
		//返回信息给客户端
		_, err = conn.Write([]byte(bufReturn))
		ErrorNotice(err)
	}
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
	res := "time:" + t + "  |ipAddr:" + ip + "  |receive byte:" + string(msgRec) + "  |msg:" + msg
	fmt.Println(res)
}
