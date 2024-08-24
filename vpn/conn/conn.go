package conn

type Conn struct{}

func NewConn() *Conn {
	return new(Conn)
}
