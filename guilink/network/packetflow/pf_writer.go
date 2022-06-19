package packetflow

import (
	"bytes"
	"errors"
	"io"

	"github.com/bitwormhole/guilink4go/guilink/dxo"
)

// Writer 包写入接口
type Writer interface {
	Write(p *Packet) error
}

////////////////////////////////////////////////////////////////////////////////

// StreamWriter 实现包写者
type StreamWriter struct {
	stream io.Writer
	closer io.Closer
}

func (inst *StreamWriter) _Impl() (Writer, io.Closer) {
	return inst, inst
}

func (inst *StreamWriter) Write(p *Packet) error {

	// size0 : 包长度
	// size1 : 头部长度
	// size2 : 主体长度

	head := []byte(p.Head)
	body := p.Body
	if body == nil {
		body = []byte{}
	}

	size1 := 4 + len(head)
	size2 := 4 + len(body)
	size0 := 4 + size1 + size2
	buffer := &bytes.Buffer{}

	inst.writeInt32(size0, buffer)
	inst.writeInt32(size1, buffer)
	buffer.Write(head)
	inst.writeInt32(size2, buffer)
	buffer.Write(body)

	return inst.writeToStream(buffer.Bytes())
}

func (inst *StreamWriter) writeToStream(src []byte) error {
	dst := inst.stream
	if dst == nil {
		return errors.New("dst==nil")
	}
	if src == nil {
		return errors.New("src==nil")
	}
	want := len(src)
	have := 0
	p := src
	for have < want {
		cb, err := dst.Write(p)
		if cb > 0 {
			have += cb
			p = p[cb:]
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *StreamWriter) writeInt32(n int, dst *bytes.Buffer) {
	buffer := [4]byte{}
	buf := buffer[:]
	i32 := dxo.Int32(n)
	i32.WriteTo(buf)
	dst.Write(buf)
}

// Close 关闭写者
func (inst *StreamWriter) Close() error {
	cl := inst.closer
	inst.stream = nil
	inst.closer = nil
	if cl != nil {
		return cl.Close()
	}
	return nil
}

// NewStreamWriter 新建写者
func NewStreamWriter(dst io.Writer, cl io.Closer) Writer {
	return &StreamWriter{stream: dst, closer: cl}
}
