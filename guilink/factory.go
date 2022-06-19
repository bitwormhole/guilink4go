package guilink

import (
	"github.com/bitwormhole/guilink4go/guilink/backend"
	"github.com/bitwormhole/guilink4go/guilink/frontend"
)

// Factory 用来创建端子
type Factory interface {
	NewClient() frontend.Client
	NewServer() backend.Server
}
