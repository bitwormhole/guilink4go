package vo

import "github.com/bitwormhole/guilink4go/guilink/dto"

// Packet 表示一个包含若干dto的包
type Packet struct {
	Base

	BoxList   []*dto.Box   `json:"boxes"`
	EventList []*dto.Event `json:"events"`
	StyleList []*dto.Style `json:"styles"`
}
