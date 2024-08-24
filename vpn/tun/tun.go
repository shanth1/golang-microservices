package tun

type Tun struct{}

func NewTun() *Tun {
	return new(Tun)
}

func (t *Tun) Read() []byte {
	return nil
}
