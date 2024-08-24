package tun

type Event int

const (
	EventUp = 1 << iota
	EventDown
	EventMTUUpdate
)

type Tun struct {
	T *NativeTun
}

func NewTun(name string, MTU int) (*Tun, error) {
	t := new(Tun)

	var err error
	if t.T, err = CreateTUN(name, MTU); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Tun) Read() []byte {
	return nil
}
