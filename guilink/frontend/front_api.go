package frontend

import "io"

// Connection 表示与会话绑定的连接
type Connection interface {
	io.Closer
}

// Client 这个接口用来操作前端的客户端
type Client interface {
	Connect(host string, port int) (Connection, error)
}
