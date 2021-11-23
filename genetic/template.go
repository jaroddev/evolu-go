package genetic

import (
	. "evolugo/chromosomes"
)

type Template interface {
	Init() Population
	Fit(*Chromosome)
	Continue(*GA) bool
}

func WithTemplates(t Template) GAOption {
	return func(g *GA) {
		g.Init = t.Init
		g.Fit = t.Fit
		g.Continue = t.Continue
	}
}
