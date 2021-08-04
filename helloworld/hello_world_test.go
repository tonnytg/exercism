package helloworld_test

import (
	"github.com/tonnytg/exercism/helloworld"
	"testing"
)

func TestSend(t *testing.T) {
	msg := helloworld.SendMsg()
	if msg != "Hello, World!" {
		t.Error("msg must be equal to Hello, World!")
	}
}
