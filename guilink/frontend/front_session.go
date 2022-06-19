package frontend

import "github.com/bitwormhole/guilink4go/guilink/vo"

// Session 表示前端的一个会话(client_connection_endpoint)
type Session struct {
	Parent     *Station
	Connection Connection

	Sender   vo.Sender
	Receiver vo.Receiver
	Handler  vo.Handler
}
