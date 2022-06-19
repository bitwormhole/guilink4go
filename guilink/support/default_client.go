package support

import (
	"errors"

	"github.com/bitwormhole/guilink4go/guilink/frontend"
)

type defaultClient struct {
	station *frontend.Station
}

func (inst *defaultClient) _Impl() frontend.Client {
	return inst
}

func (inst *defaultClient) Connect(host string, port int) (frontend.Connection, error) {
	return nil, errors.New("no impl")
}
