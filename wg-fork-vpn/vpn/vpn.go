package vpn

import (
	"encoding/hex"
	"fmt"

	"github.com/shanth1/golang-microservices/vpn/conn"
	"github.com/shanth1/golang-microservices/vpn/tun"
)

type Config struct {
}

type Vpn struct {
	T   *tun.Tun
	C   *conn.Conn
	Cfg Config
}

func NewVpn() (*Vpn, error) {
	var err error

	v := new(Vpn)
	v.C = conn.NewConn()
	v.T, err = tun.NewTun("utun5", 1500)

	if err != nil {
		return nil, err
	}
	return v, nil
}

func (v *Vpn) Run() error {
	for {
		buf, err := v.T.Read()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("> %s\n", hex.EncodeToString(buf))
		}
	}
}

func (v *Vpn) Close() {
	v.T.T.Close()
}
