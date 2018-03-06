package execkit

import (
	"testing"

	"github.com/BaritoLog/go-boilerplate/errkit"
	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestPrint(t *testing.T) {
	printer := &DummyPrinter{}
	cmds := []Cmd{
		DummyCmd{OutputBytes: []byte("hello")},
		DummyCmd{OutputBytes: []byte("world")},
	}

	err := Print(printer, cmds...)

	FatalIfError(t, err)
	FatalIf(t, len(printer.Lines) != 2, "output lines should be 2")
	FatalIf(t, printer.Lines[0] != "hello", "first output should be \"hello\"")
	FatalIf(t, printer.Lines[1] != "world", "second output should be \"hello\"")
}

func TestPrint_Error(t *testing.T) {
	printer := &DummyPrinter{}
	cmds := []Cmd{
		DummyCmd{OutputBytes: []byte("hello")},
		DummyCmd{OutputError: errkit.Error("some error")},
		DummyCmd{OutputBytes: []byte("world")},
	}

	err := Print(printer, cmds...)

	FatalIfWrongError(t, err, "some error")
	FatalIf(t, len(printer.Lines) != 1, "output lines should be 1")
}
