package main

import "github.com/bitwormhole/guilink4go"

func main() {

	server := guilink4go.GetFactory().NewServer()

	server.Start()
	server.Join()
}
