package execkit

import (
	"bytes"
	"testing"

	"github.com/BaritoLog/go-boilerplate/errkit"
	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestPrint(t *testing.T) {
	var buf bytes.Buffer

	cmds := []Cmd{
		DummyCmd{OutputBytes: []byte("hello")},
		DummyCmd{OutputBytes: []byte("world")},
	}

	err := Print(&buf, cmds...)

	want := `hello
world
`

	FatalIfError(t, err)
	FatalIf(t, buf.String() != want, "print wrong value")

}

func TestPrint_Error(t *testing.T) {
	var buf bytes.Buffer

	cmds := []Cmd{
		DummyCmd{OutputBytes: []byte("hello")},
		DummyCmd{OutputError: errkit.Error("some error")},
		DummyCmd{OutputBytes: []byte("world")},
	}

	err := Print(&buf, cmds...)

	want := `hello
`

	FatalIfWrongError(t, err, "some error")
	FatalIf(t, buf.String() != want, "print wrong value")
}
