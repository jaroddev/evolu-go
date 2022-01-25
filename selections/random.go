package selections

import (
	. "evolugo/chromosomes"
	"math/rand"
)

type Random struct {
	ParentNumber int
}

// How can we test random behaviour ???
func (r *Random) Select(pop *Population) Population {
	if r.ParentNumber > len((*pop)) {
		panic("Not Enough population element")
	}

	parents := make(Population, 0, r.ParentNumber)

	picks := rand.Perm(len(*pop))
	picks = picks[:r.ParentNumber]

	for _, pick := range picks {
		parents = append(parents, (*pop)[pick])
	}

	return parents
}
