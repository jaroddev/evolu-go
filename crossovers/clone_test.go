package crossovers

import (
	. "evolugo/chromosomes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cloneTwo  Clone = Clone{ChildrenNumber: 2}
	cloneFour Clone = Clone{ChildrenNumber: 4}

	parents Population = Population{
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
)

func TestCloneTwo(t *testing.T) {
	children := cloneTwo.Cross(&parents)
	assert.Equal(t, len(children), 2)
}

func TestTryToCloneTooManyChromosome(t *testing.T) {
	children := cloneFour.Cross(&parents)
	assert.NotEqual(t, len(children), 4)
}
