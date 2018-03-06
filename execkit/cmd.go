package execkit

type Cmd interface {
	Output() ([]byte, error)
}

type DummyCmd struct {
	OutputBytes []byte
	OutputError error
}

func (c DummyCmd) Output() ([]byte, error) {
	return c.OutputBytes, c.OutputError
}
