package selections

import (
	. "evolugo/chromosomes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getLargeTestPopulation() Population {
	return Population{
		Chromosome{
			Age:     0,
			Fitness: 15,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 250,
			Alleles: []bool{
				true, false, false, true, false, true, false, false,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 275,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 25,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 150,
			Alleles: []bool{
				true, false, false, true, false, true, false, false,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 1750,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 25,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 3,
			Alleles: []bool{
				true, false, false, true, false, true, false, false,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 1,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
	}
}

func TestTwoChromosomeWereReturned(t *testing.T) {
	pop := getLargeTestPopulation()
	selection := Tournament{
		Participant: 5,
		Winner:      2,
	}

	parents := selection.Select(&pop)
	assert.Equal(t, len(parents), 2)
}
