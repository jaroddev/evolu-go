package selections

import (
	. "evolugo/chromosomes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	selectTwo SelectBest = SelectBest{
		ParentNumber: 2,
	}
	selectFour SelectBest = SelectBest{
		ParentNumber: 4,
	}

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

func TestSelectTwo(t *testing.T) {
	children := selectTwo.Select(&parents)
	assert.Equal(t, len(children), 2)
}

func TestSelectTooMany(t *testing.T) {
	children := selectFour.Select(&parents)
	assert.NotEqual(t, len(children), 4)
}
