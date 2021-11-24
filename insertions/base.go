package insertions

import (
	. "evolugo/chromosomes"
)

type Insertion func(Population, Population) Population

func AddChildren(pop Population, children Population) Population {
	pop = append(pop, children...)
	return pop
}
