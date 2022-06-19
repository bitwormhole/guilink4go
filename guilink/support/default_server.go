package support

import (
	"errors"
	"io"
	"net"
	"strconv"
	"time"

	"log"

	"github.com/bitwormhole/guilink4go/guilink"
	"github.com/bitwormhole/guilink4go/guilink/backend"
	"github.com/bitwormhole/guilink4go/guilink/network/packetflow"
	"github.com/bitwormhole/guilink4go/guilink/vo"
)

type defaultServer struct {
	station *backend.Station
	runtime *defaultServerRuntime

	starting bool
	stopping bool
}

func (inst *defaultServer) _Impl() backend.Server {
	return inst
}

func (inst *defaultServer) Start() error {

	if inst.runtime != nil || inst.starting {
		return errors.New("bad status")
	}

	rt := &defaultServerRuntime{}
	rt.parent = inst
	inst.runtime = rt
	inst.starting = true
	go func() {
		rt.run()
	}()
	return nil
}

func (inst *defaultServer) Stop() error {
	inst.stopping = true
	rt := inst.runtime
	if rt != nil {
		return rt.close()
	}
	return nil
}

func (inst *defaultServer) Join() error {
	rt := inst.runtime
	if rt == nil {
		return errors.New("bad status")
	}
	for {
		if rt.stopped {
			break
		}
		time.Sleep(time.Second * 5)
	}
	return nil
}

func (inst *defaultServer) IsRunning() bool {
	rt := inst.runtime
	if rt != nil {
		return !rt.stopped
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////

type defaultServerRuntime struct {
	parent      *defaultServer
	netListener net.Listener

	started bool
	stopped bool
}

func (inst *defaultServerRuntime) run() {

	defer func() {
		x := recover()
		inst.handleErrorEx(x)
		inst.stopped = true
	}()

	err := inst.run2()
	if err != nil {
		inst.handleError(err)
	}
}

func (inst *defaultServerRuntime) handleError(err error) {
	if err == nil {
		return
	}
	log.Println(err)
}

func (inst *defaultServerRuntime) handleErrorEx(x any) {
	if x == nil {
		return
	}
	err, ok := x.(error)
	if ok {
		inst.handleError(err)
	} else {
		log.Println("error: ", x)
	}
}

func (inst *defaultServerRuntime) run2() error {

	const tcp = "tcp"
	station := inst.parent.station
	host := station.Host
	port := station.Port
	if port < 1 {
		port = guilink.DefaultPort
	}
	address := host + ":" + strconv.Itoa(port)

	addr, err := net.ResolveTCPAddr(tcp, address)
	if err != nil {
		return err
	}

	li, err := net.ListenTCP(tcp, addr)
	if err != nil {
		return err
	}

	defer func() {
		err = li.Close()
		inst.handleError(err)
	}()

	inst.netListener = li
	inst.started = true
	log.Println("Listen Guilink(over_TCP) at ", li.Addr())

	for {
		conn, err := li.AcceptTCP()
		if err != nil {
			inst.handleError(err)
			time.Sleep(time.Second)
			continue
		}
		err = inst.initSession(conn)
		if err != nil {
			inst.handleError(err)
			err = conn.Close()
			inst.handleError(err)
			time.Sleep(time.Second)
			continue
		}
		if inst.parent.stopping {
			break
		}
	}
	return nil
}

func (inst *defaultServerRuntime) initSession(conn *net.TCPConn) error {

	session := &defaultServerSession{}
	sender := &defaultServerSender{}

	session.conn = conn
	session.inner.Parent = inst.parent.station
	session.inner.Sender = sender.init(conn)
	session.inner.Handler = nil // todo ...

	// start session
	go func() {
		err := session.run()
		inst.handleError(err)
	}()
	return nil
}

func (inst *defaultServerRuntime) close() error {
	li := inst.netListener
	inst.netListener = nil
	if li != nil {
		return li.Close()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type defaultServerSession struct {
	inner   backend.Session
	runtime *defaultServerRuntime
	conn    net.Conn

	started bool
	stopped bool
}

func (inst *defaultServerSession) run() error {
	conn := inst.conn
	if conn == nil {
		return errors.New("conn==nil")
	}
	defer func() {
		inst.stopped = true
		err := conn.Close()
		inst.runtime.handleError(err)
	}()
	inst.started = true
	return inst.runRxLoop(conn)
}

func (inst *defaultServerSession) runRxLoop(conn net.Conn) error {
	reader := packetflow.NewStreamReader(conn, nil)
	handler := inst.inner.Handler
	for !inst.isClosed() {
		p1, err := reader.Read()
		if err != nil {
			return err
		}
		p2, err := decodePacket(p1)
		err = handler.Handle(p2)
		if err != nil {
			inst.runtime.handleError(err)
		}
	}
	return nil
}

func (inst *defaultServerSession) isClosed() bool {
	conn := inst.conn
	if conn == nil {
		return true
	}
	if inst.runtime.parent.stopping {
		return true
	}
	return false
}

func (inst *defaultServerSession) close() error {
	conn := inst.conn
	inst.conn = nil
	if conn == nil {
		return nil
	}
	return conn.Close()
}

////////////////////////////////////////////////////////////////////////////////

type defaultServerSender struct {
	writer packetflow.Writer
}

func (inst *defaultServerSender) init(w io.Writer) vo.Sender {
	inst.writer = packetflow.NewStreamWriter(w, nil)
	return nil
}

// func (inst *defaultServerSender) close() error {
// }

func (inst *defaultServerSender) Send(p1 *vo.Packet) error {
	p2, err := encodePacket(p1)
	if err != nil {
		return err
	}
	return inst.writer.Write(p2)
}
