package selections

import (
	. "evolugo/chromosomes"
)

type Tournament struct {
	Participant int
	Winner      int
}

func (t *Tournament) Select(pop *Population) Population {
	// Randomly
	random := Random{
		ParentNumber: t.Participant,
	}
	selection := SelectBest{
		ParentNumber: t.Winner,
	}

	participant := random.Select(pop)
	return selection.Select(&participant)
}
