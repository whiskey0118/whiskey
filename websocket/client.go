package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

type Client struct {
	Device    *Device
	Conn      *websocket.Conn
	Pool      *Pool
	InChan    chan *Message
	OutChan   chan *Message
	Done      chan bool
	Interrupt chan os.Signal //捕捉客户端进程关闭信号
	Token     string
}

func InitClient(conn *websocket.Conn, device *Device, pool *Pool, token string) (client *Client, err error) {
	client = &Client{
		Device:  device,
		Conn:    conn,
		Pool:    pool,
		InChan:  make(chan *Message, 100),
		OutChan: make(chan *Message, 100),
		Token:   token,
	}
	return
}

func (client *Client) ReadMessage() {
	for {
		select {
		case message := <-client.InChan:
			//刚开始因为下面这行总是报错：panic: runtime error: invalid memory address or nil pointer dereference，原来自己的client.Device.Uid 没有赋值，还是nil -。-
			//log.Printf("Read message uid:%s message:%s", client.Device.Uid, message.Body)
			log.Printf("Read message type:%v message:%v", message.Type, message.Body)
		case <-client.Done:
			client.Close()
			return
		case <-client.Interrupt:
			log.Println("interrupt")
			err := client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, fmt.Sprintf("%s closed connection", client.Device.Uid)))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			client.Close()
			return
		}
	}
}

func (client *Client) WriteMessage() {
	var message Message
	for {
		_, b, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Println("read body err:", err)
			break
		}
		err = json.Unmarshal(b, &message)
		if err != nil {
			log.Println(err)
			return
		}
		client.InChan <- &message
	}
}

func (client *Client) Broadcast(message *Message) {
	for {
		select {
		case client.OutChan <- message:
			log.Printf("Write message to uid:%s message:%s", client.Device.Uid, message.Body)
		case <-client.Interrupt:
			log.Println("interrupt")
			err := client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, fmt.Sprintf("%s closed connection", client.Device.Uid)))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			client.Close()
		}
	}
}

func (client *Client) Send() {

}

func (client *Client) Close() {
	client.Conn.Close()
}
