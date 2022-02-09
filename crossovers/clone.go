package crossovers

import (
	. "github.com/jaroddev/evolugo/chromosomes"
)

type Clone struct {
	ChildrenNumber int
}

// Should be renamed Disabled !!!
func (c *Clone) Cross(pop *Population) Population {

	if c.ChildrenNumber > len((*pop)) || c.ChildrenNumber > cap((*pop)) {
		return (*pop)[:len((*pop))]
	}

	return (*pop)[:c.ChildrenNumber]

}
