package vo

type Receiver interface {
	Receive() (*Packet, error)
}

type Handler interface {
	Handle(p *Packet) error
}
