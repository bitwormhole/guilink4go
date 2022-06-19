package vo

type Sender interface {
	Send(o *Packet) error
}
