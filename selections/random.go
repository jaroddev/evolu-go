package selections

import (
	. "evolugo/chromosomes"
	"math/rand"
)

// How can we test random behaviour ???
func RandomSelection(parentNumber int) Selection {
	return func(pop *Population) Population {
		if parentNumber > len((*pop)) {
			panic("Not Enough population element")
		}

		parents := make(Population, 0, parentNumber)

		picks := rand.Perm(len(*pop))
		picks = picks[:parentNumber]

		for _, pick := range picks {
			parents = append(parents, (*pop)[pick])
		}

		return parents
	}
}
