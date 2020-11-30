package websocket

import (
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
	Done      chan int
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
	//tic := time.NewTicker(2 * time.Second)
	//defer tic.Stop()

	for {
		select {
		case message := <-client.InChan:
			log.Printf("Read message uid:%s message:%s", client.Device.Uid, message.Body)
		case <-client.Done:
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

func (client *Client) WriteMessage(message *Message) {
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

func (client *Client) Close() {
	client.Conn.Close()
}
