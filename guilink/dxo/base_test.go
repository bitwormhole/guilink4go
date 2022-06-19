package dxo

import (
	"strings"
	"testing"
)

func TestInt32(t *testing.T) {
	values := []Int32{-8, -4, -2, -1, 0, 1, 2, 4, 8}
	for n := Int32(8); n > 0; n <<= 1 {
		values = append(values, n)
		values = append(values, -n)
	}
	buffer := []byte{0, 0, 0, 0}
	for _, n1 := range values {
		err := n1.WriteTo(buffer)
		if err != nil {
			panic(err)
		}
		n2, err := ReadInt32From(buffer)
		if err != nil {
			panic(err)
		}
		msg := strings.Builder{}
		if n1 == n2 {
			msg.WriteString("test Int32(")
			msg.WriteString(n1.String())
			msg.WriteString(")...ok")
			t.Log(msg.String())
		} else {
			msg.WriteString("test Int32, want ")
			msg.WriteString(n1.String())
			msg.WriteString(" have ")
			msg.WriteString(n2.String())
			panic(msg.String())
		}
	}
	t.Log("OK")
}
