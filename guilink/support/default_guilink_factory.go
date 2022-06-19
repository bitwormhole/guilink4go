package support

import (
	"github.com/bitwormhole/guilink4go/guilink"
	"github.com/bitwormhole/guilink4go/guilink/backend"
	"github.com/bitwormhole/guilink4go/guilink/frontend"
)

// DefaultGuilinkFactory 实现默认的 guilink 工厂
type DefaultGuilinkFactory struct{}

func (inst *DefaultGuilinkFactory) _Impl() guilink.Factory {
	return inst
}

// NewClient 创建客户端
func (inst *DefaultGuilinkFactory) NewClient() frontend.Client {

	station := &frontend.Station{}
	client := &defaultClient{}

	// wire
	client.station = station
	station.Client = client

	return station.Client
}

// NewServer 创建服务端
func (inst *DefaultGuilinkFactory) NewServer() backend.Server {

	station := &backend.Station{}
	server := &defaultServer{}

	// wire
	server.station = station
	station.Server = server

	return station.Server
}
