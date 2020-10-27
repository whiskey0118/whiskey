package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"whiskey/proxy"
	"whiskey/proxy/direct"
	"whiskey/tools"
)

var (
	version    = "v0.1.0"
	configFile = flag.String("configFile", "config.json", "config file name")
)

type Config struct {
	Local  string `json:"local"`
	Route  string `json:"route"`
	Remote string `json:"remote"`
}

func loadConfig(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	config := Config{}

	buf := make([]byte,1024)
	n,err := f.Read(buf)
	err = json.Unmarshal(buf[:n],&config)
	if err != nil {
		return nil,err
	}
	return &config,nil
}

func main() {
	flag.Parse()
	path := tools.GetPath(*configFile)
	config,err:=loadConfig(path)
	if err != nil {
		fmt.Println(err)
	}

	//创建一个服务端
	localServer,err :=proxy.ServerFromURL(config.Local)
	if err != nil {
		log.Fatal("create server err:",err)
		os.Exit(1)
	}
	defer localServer.Stop()

	//创建客户端
	remoteClient,err := proxy.ClientFromURL(config.Remote)
	if err != nil {
		log.Fatal("create client err:",err)
		os.Exit(1)
	}

	//directClient, _ := proxy.ClientFromURL("direct://")

	//侦听本地端口
	listener,err := net.Listen("tcp",localServer.Addr())
	if err != nil {
		log.Fatal("listen local port err :",localServer.Addr(),err)
		os.Exit(1)
	}
	log.Printf("%v listening TCP on %v",localServer.Name(),localServer.Addr())
	//创建协程接受连接
	go func() {
		for {
			conn,err := listener.Accept()
			if err != nil {
				log.Fatal(err)
				break
			}

			//处理链接
			go func() {
				defer conn.Close()
				var client proxy.Client
				wconn,targetAddr,err := localServer.Handshake(conn)
				if err != nil {
					log.Fatal("failed in handshake:",err)
					return
				}

				client = remoteClient
				//连接客户端地址
				clientAddr := remoteClient.Addr()
				if _,ok := client.(*direct.Direct);ok{
					clientAddr = targetAddr.String()
				}
				rc,err := net.Dial("tcp",clientAddr)
				if err != nil {
					log.Fatal("dial client err :",err)
					return
				}
				defer rc.Close()

				// 不同的客户端协议各自实现自己的请求逻辑
				wrc, err := client.Handshake(rc, targetAddr.String())
				if err != nil {
					log.Printf("failed in handshake to %v: %v", clientAddr, err)
					return
				}
				// 流量转发
				go io.Copy(wrc, wconn)
				io.Copy(wconn, wrc)

			}()

		}
	}()

	{
		osSignals := make(chan os.Signal,1)
		signal.Notify(osSignals,os.Interrupt,os.Kill,syscall.SIGTERM)
		<-osSignals
	}
}


