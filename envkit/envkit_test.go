package envkit

import (
	"os"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestGetString(t *testing.T) {
	os.Setenv("some-key", "some-value")
	s := GetString("some-key", "default-value")

	FatalIf(t, s != "some-value", "wrong return")
}

func TestGetString_WrongKey(t *testing.T) {
	s := GetString("wrong-key", "default-value")
	FatalIf(t, s != "default-value", "wrong return")
}

func TestGetInt_WrongKey(t *testing.T) {
	i := GetInt("wrong-key", 9999)
	FatalIf(t, i != 9999, "wrong return")
}

func TestGetInt(t *testing.T) {
	os.Setenv("some-key", "8888")
	i := GetInt("some-key", 9999)

	FatalIf(t, i != 8888, "wrong return")
}

func TestGetInt_NaN(t *testing.T) {

	os.Setenv("some-key", "nan")
	i := GetInt("some-key", 9999)

	FatalIf(t, i != 9999, "wrong return")
}
