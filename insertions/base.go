package insertions

import (
	. "github.com/jaroddev/evolugo/chromosomes"
)

type Add struct{}

func (a *Add) Insert(pop Population, children Population) Population {
	pop = append(pop, children...)
	return pop
}
