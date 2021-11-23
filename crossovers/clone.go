package crossovers

import (
	. "evolugo/chromosomes"
)

type CrossOver func(p *Population) Population

func CloneFirst(childrenNumber int) CrossOver {
	return func(p *Population) Population {
		if childrenNumber > len((*p)) || childrenNumber > cap((*p)) {
			return (*p)[:len((*p))]
		}

		return (*p)[:childrenNumber]
	}
}
