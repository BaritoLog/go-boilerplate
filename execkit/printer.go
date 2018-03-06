package execkit

type Printer interface {
	Print(s string)
}

type DummyPrinter struct {
	Lines []string
}

func (p *DummyPrinter) Print(s string) {
	p.Lines = append(p.Lines, s)
}
