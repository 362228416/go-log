package log

import "testing"

func TestInfo(t *testing.T) {
	Debug("hi")
	Info("msg")
	Warning("fbi")
	Error("wtf")
}
