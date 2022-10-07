package log

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	os.Setenv("STDLOG", "0")
	lg1 := NewLogger("lg1", "../logs")
	lg1.Debug("from lg1")
	lg1.Info("from lg1")

	lg2 := NewLogger("lg2")
	lg2.Debug("from lg2")
	lg2.Info("from lg2")

	lg3 := NewLogger("lg1")
	lg3.Debug("from lg3")
	lg3.Info("from lg3")
}
