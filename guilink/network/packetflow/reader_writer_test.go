package packetflow

import (
	"bytes"
	"errors"
	"testing"
)

func TestReaderAndWriter(t *testing.T) {

	pktInit := func(head string, body string) *Packet {
		p := &Packet{}
		p.Head = head
		if len(body) > 0 {
			p.Body = []byte(body)
		}
		return p
	}

	pktEq := func(p1, p2 *Packet) bool {
		n := bytes.Compare(p1.Body, p2.Body)
		return (n == 0) && (p1.Head == p2.Head)
	}

	list := []*Packet{}
	list = append(list, pktInit("", ""))
	list = append(list, pktInit("", "uvwxyz"))
	list = append(list, pktInit("abcdefg", ""))
	list = append(list, pktInit("abcdefg", "uvwxyz"))

	fifo := &bytes.Buffer{}
	writer := NewStreamWriter(fifo, nil)

	for _, pk := range list {
		err := writer.Write(pk)
		if err != nil {
			panic(err)
		}
	}

	reader := NewStreamReader(fifo, nil)

	for i := 0; i < len(list); {
		p1 := list[i]
		p2, err := reader.Read()
		if err != nil {
			panic(err)
		}
		if pktEq(p1, p2) {
			i++
		} else {
			panic(errors.New("p1 != p2"))
		}
	}

	t.Log("OK")
}
