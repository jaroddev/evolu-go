package selections

import (
	. "evolugo/chromosomes"
)

type SelectBest struct {
	filter       GetBest
	ParentNumber int
}

func (selection *SelectBest) Select(p *Population) Population {
	if selection.ParentNumber > len((*p)) || selection.ParentNumber > cap((*p)) {
		return (*p)[:len((*p))]
	}

	return FilterWithLimit(*p, &selection.filter, selection.ParentNumber)
}
