package selections

import (
	. "evolugo/chromosomes"
)

func Tournament(tournamentParticipant, winner int) Selection {

	return func(pop *Population) Population {
		// Randomly
		tournament := RandomSelection(tournamentParticipant)
		best := SelectFirst(winner)

		member := tournament(pop)
		return best(&member)
	}
}
