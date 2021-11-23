package selections

import (
	. "evolugo/chromosomes"
)

type Selection func(p *Population) Population

func SelectFirst(parentNumber int) Selection {
	return func(p *Population) Population {
		if parentNumber > len((*p)) || parentNumber > cap((*p)) {
			return (*p)[:len((*p))]
		}

		return (*p)[:parentNumber]
	}
}
