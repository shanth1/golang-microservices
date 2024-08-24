package conn

import (
	"encoding/hex"
	"fmt"
)

// Config of http server
type Config struct {
}

// Conn Websocket connection
type Conn struct {
	Cfg   *Config
	Peers map[string]string
}

func NewConn() *Conn {
	return new(Conn)
}

func (c *Conn) Write(data []byte) {
	fmt.Printf("> %s\n", hex.EncodeToString(data))
}
