package packetflow

import "github.com/bitwormhole/guilink4go/guilink/dxo"

// Packet 表示一个pktline包
type Packet struct {
	Size dxo.Int32 // 32 位包长度（包含'Size'字段的4个字节）
	Head string    // 文本形式的头部
	Body []byte    // 二进制形式的主体
}
