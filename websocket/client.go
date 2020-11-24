package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"time"
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

func (client *Client) ReadMessage() (mess *Message, err error) {
	tic := time.NewTicker(2 * time.Second)
	defer tic.Stop()

	for {
		select {
		case <-tic.C:
			mess.Body = time.Now().String()
			mess.Type = client.Device.Uid
			res, _ := json.Marshal(&mess)
			err = client.Conn.WriteMessage(websocket.TextMessage, res)
			if err != nil {
				log.Printf("%s write err ", client.Device.Uid)
			}
		case <-client.Done:
			return
		case <-client.Interrupt:
			log.Println("interrupt")
			err := client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, fmt.Sprintf("%s closed connection", client.Device.Uid)))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}

func (client *Client) WriteMessage(message *Message) (err error) {
	client.OutChan <- message
	return
}

func (client *Client) Close() {
	client.Conn.Close()
}
