package guilink4go

import (
	"github.com/bitwormhole/guilink4go/guilink"
	"github.com/bitwormhole/guilink4go/guilink/support"
)

const (
	theModuleName = "github.com/bitwormhole/guilink4go"
	theModuleVer  = "v0.0.1"
	theModuleRev  = 1
)

// go:embed "src/main/resources"
// var theModuleMainRes embed.FS

// GetFactory 取工厂
func GetFactory() guilink.Factory {
	return &support.DefaultGuilinkFactory{}
}
