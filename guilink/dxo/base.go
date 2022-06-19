package dxo

import (
	"errors"
	"strconv"
)

// Int32 表示一个32位有符号整数
type Int32 int32

func (v Int32) String() string {
	n := int64(v)
	return strconv.FormatInt(n, 10)
}

// Int 把值转换成int格式
func (v Int32) Int() int {
	return int(v)
}

// WriteTo 把值写入缓冲区 (按网络字节序,Big Endian)
func (v Int32) WriteTo(buffer []byte) error {
	if buffer == nil {
		return errors.New("buffer is nil")
	}
	const wantSize = 4
	size := len(buffer)
	if size < wantSize {
		return errors.New("bad Int32 size")
	}
	n := v
	for i := wantSize - 1; i >= 0; i-- {
		buffer[i] = byte(0xff & n)
		n >>= 8
	}
	return nil
}

// ReadInt32From 从缓冲区读取 Int32 值 (按网络字节序, Big Endian)
func ReadInt32From(buffer []byte) (Int32, error) {
	if buffer == nil {
		return 0, errors.New("buffer is nil")
	}
	const wantSize = 4
	size := len(buffer)
	if size < wantSize {
		return 0, errors.New("bad Int32 size")
	}
	n := Int32(0)
	for i := 0; i < wantSize; i++ {
		b := buffer[i]
		n <<= 8
		n |= (0xff & Int32(b))
	}
	return n, nil
}
