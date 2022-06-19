package network

import (
	"crypto/tls"
	"errors"
	"log"
	"net"
	"strconv"
	"time"
)

////////////////////////////////////////////////////////////////////////////////

type StreamService interface {
	Handle(conn net.Conn) error
}

type StreamServiceFactory interface {
	Create() StreamService
}

////////////////////////////////////////////////////////////////////////////////

// TLSServer 封装一个简单的TLS服务器
type TLSServer struct {
	Port            int
	PrivateKeyFile  string
	CertificateFile string
	ServiceFactory  StreamServiceFactory

	// status
	starting bool
	stopping bool
	started  bool
	stopped  bool
}

// Start 启动服务
func (inst *TLSServer) Start() error {
	if inst.starting || inst.stopping {
		return errors.New("bad status")
	}
	inst.starting = true
	go func() {
		err := inst.run()
		if err != nil {

		}
	}()
	return nil
}

// Stop 停止服务
func (inst *TLSServer) Stop() error {
	inst.stopping = true
	return nil
}

// Join 等待服务停止
func (inst *TLSServer) Join() error {
	for {
		if inst.starting && !inst.stopped {
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}
	return nil
}

func (inst *TLSServer) handleError(err error) {
	log.Println(err)
}

func (inst *TLSServer) handleConn(conn net.Conn) {
	ser := inst.ServiceFactory.Create()
	err := ser.Handle(conn)
	if err != nil {
		inst.handleError(err)
	}
}

func (inst *TLSServer) run() error {

	defer func() {
		inst.stopped = true
	}()

	cert, err := tls.LoadX509KeyPair(inst.CertificateFile, inst.PrivateKeyFile)
	if err != nil {
		return err
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	addr := ":" + strconv.Itoa(inst.Port)
	listener, err := tls.Listen("tcp", addr, cfg)
	if err != nil {
		return err
	}

	defer listener.Close()
	inst.started = true
	addrStr := listener.Addr().String()
	log.Println("Listen [", addrStr, "] over TLS")

	for {
		if inst.stopping {
			break
		}
		conn, err := listener.Accept()
		if err == nil {
			go inst.handleConn(conn)
		} else {
			inst.handleError(err)
		}
	}

	return nil
}
