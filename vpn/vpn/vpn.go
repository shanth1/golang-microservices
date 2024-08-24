package vpn

import (
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
	v := new(Vpn)
	v.C = conn.NewConn()
	v.T = tun.NewTun()

	return v, nil
}

func (v Vpn) Run() error {
	for {
		v.T.Read()
	}
}
