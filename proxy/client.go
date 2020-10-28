package proxy

import (
	"io"
	"net"
	"net/url"
)

const name = "direct"

func init() {
	RegisterClient(name, NewDirectClient)
}

func NewDirectClient(url *url.URL) (Client, error) {
	return &Direct{}, nil
}

type Direct struct{}

func (d *Direct) Name() string { return name }

func (d *Direct) Addr() string { return name }

func (d *Direct) Handshake(underlay net.Conn, target string) (io.ReadWriter, error) {
	return underlay, nil
}
