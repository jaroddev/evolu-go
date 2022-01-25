package selections

import (
	. "evolugo/chromosomes"
)

type SelectFirst struct {
	ParentNumber int
}

func (selection *SelectFirst) Select(p *Population) Population {
	if selection.ParentNumber > len((*p)) || selection.ParentNumber > cap((*p)) {
		return (*p)[:len((*p))]
	}

	return (*p)[:selection.ParentNumber]
}
