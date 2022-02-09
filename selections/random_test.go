package selections

import (
	"testing"

	. "github.com/jaroddev/evolugo/chromosomes"

	"github.com/stretchr/testify/assert"
)

var (
	randomSelectTwo Random = Random{
		ParentNumber: 2,
	}
	randomSelectFour Random = Random{
		ParentNumber: 4,
	}
)

func getTestPopulation() Population {
	return Population{
		Chromosome{
			Age:     0,
			Fitness: 0,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 0,
			Alleles: []bool{
				true, false, false, true, false, true, false, false,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 0,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
	}
}

func TestValidRandomSelectionLength(t *testing.T) {
	pop := getTestPopulation()
	parentsWithTwoChromosome := randomSelectTwo.Select(&pop)

	assert.Equal(t, len(parentsWithTwoChromosome), 2)
}

func TestNotValidRandomSelectionLength(t *testing.T) {
	pop := getTestPopulation()

	assert.Panics(t, func() {
		randomSelectFour.Select(&pop)
	})

}
