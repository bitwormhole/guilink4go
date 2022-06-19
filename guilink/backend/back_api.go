package backend

import "io"

// Connection 表示与会话绑定的连接
type Connection interface {
	io.Closer
}

// Server 这个接口用来控制后端的服务器
type Server interface {
	Start() error
	Stop() error
	Join() error
	IsRunning() bool
}
