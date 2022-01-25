package crossovers

import (
	. "evolugo/chromosomes"
)

type Clone struct {
	ChildrenNumber int
}

// Should be cloning best only ???
// Is it used as a way to disable crossovers
func (c *Clone) Cross(pop *Population) Population {

	if c.ChildrenNumber > len((*pop)) || c.ChildrenNumber > cap((*pop)) {
		return (*pop)[:len((*pop))]
	}

	return (*pop)[:c.ChildrenNumber]

}
