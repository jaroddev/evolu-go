package selections

import (
	. "evolugo/chromosomes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	randomSelectTwo  Selection = RandomSelection(2)
	randomSelectFour Selection = RandomSelection(4)
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
	parentsWithTwoChromosome := randomSelectTwo(&pop)

	assert.Equal(t, len(parentsWithTwoChromosome), 2)
}

func TestNotValidRandomSelectionLength(t *testing.T) {
	pop := getTestPopulation()

	assert.Panics(t, func() {
		randomSelectFour(&pop)
	})

}
