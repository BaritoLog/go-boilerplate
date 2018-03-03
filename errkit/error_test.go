package errkit

import (
	"testing"

	"github.com/imantung/go-boilerplate/testkit"
)

func TestError(t *testing.T) {
	err := Error("error1")
	testkit.FatalIfWrongError(t, err, "error1")
}
