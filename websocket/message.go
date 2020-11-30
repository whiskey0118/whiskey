package websocket

type Message struct {
	//Time time.Time `json:"time"`
	Type string `json:"type"`
	Body string `json:"body"`
}
