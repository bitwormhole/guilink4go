package packetflow

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/bitwormhole/guilink4go/guilink/dxo"
)

// Reader 包读取接口
type Reader interface {
	Read() (*Packet, error)
}

////////////////////////////////////////////////////////////////////////////////

// StreamReader 实现读取接口
type StreamReader struct {
	source io.Reader
	closer io.Closer
}

func (inst *StreamReader) _Impl() (Reader, io.Closer) {
	return inst, inst
}

func (inst *StreamReader) Read() (*Packet, error) {

	// size0 : 包长度
	// size1 : 头部长度
	// size2 : 主体长度

	// read size
	int32buffer := [4]byte{}
	size0, err := inst.readInt32(int32buffer[:])
	if err != nil {
		return nil, err
	}

	// read raw data
	size1plus2 := size0 - 4
	rawbuffer := make([]byte, size1plus2)
	err = inst.readStream(rawbuffer, size1plus2)
	if err != nil {
		return nil, err
	}

	// get head
	size1, head, err := inst.getPartFrom(rawbuffer)
	if err != nil {
		return nil, err
	}

	// get body
	size2, body, err := inst.getPartFrom(rawbuffer[size1:])
	if err != nil {
		return nil, err
	}

	// check size
	if size1+size2 != size1plus2 {
		return nil, errors.New("bad packet size")
	}

	p := &Packet{}
	p.Size = size0
	p.Head = string(head)
	p.Body = body
	return p, nil
}

func (inst *StreamReader) getPartFrom(b []byte) (dxo.Int32, []byte, error) {
	size, err := dxo.ReadInt32From(b)
	if err != nil {
		return 0, nil, err
	}
	want := size.Int()
	have := len(b)
	datalen := want - 4
	if have < want || datalen < 0 {
		msg := strings.Builder{}
		msg.WriteString("bad size, want ")
		msg.WriteString(strconv.Itoa(want))
		msg.WriteString(", but have ")
		msg.WriteString(strconv.Itoa(have))
		return 0, nil, errors.New(msg.String())
	}
	data := b[4:want]
	return size, data, nil //  (length,data,err)
}

func (inst *StreamReader) readInt32(b []byte) (dxo.Int32, error) {
	const wantSize = 4
	err := inst.readStream(b, wantSize)
	if err != nil {
		return 0, err
	}
	return dxo.ReadInt32From(b)
}

func (inst *StreamReader) readStream(b []byte, wantSize dxo.Int32) error {
	src := inst.source
	if src == nil {
		return errors.New("no source stream or closed")
	}
	want := wantSize.Int()
	have := 0
	p := b[0:want]
	for have < want {
		cb, err := src.Read(p)
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

// Close 关闭读者
func (inst *StreamReader) Close() error {
	cl := inst.closer
	inst.source = nil
	inst.closer = nil
	if cl != nil {
		return cl.Close()
	}
	return nil
}

// NewStreamReader 新建读者
func NewStreamReader(src io.Reader, closer io.Closer) Reader {
	return &StreamReader{source: src, closer: closer}
}
