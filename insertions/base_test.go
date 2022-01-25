package insertions

import (
	. "evolugo/chromosomes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getBaseParents() Population {
	return Population{
		Chromosome{
			Age:     0,
			Fitness: 10,
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

func getBaseChildren() Population {
	return Population{
		Chromosome{
			Age:     0,
			Fitness: 15,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 7,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		},
	}
}

func TestAddChildren(t *testing.T) {
	parents := getBaseParents()
	children := getBaseChildren()

	insertion := Add{}

	newParents := insertion.Insert(parents, children)
	assert.Equal(t, len(newParents), len(parents)+len(children))
}

func TestTwoAddChildren(t *testing.T) {
	parents := getBaseParents()
	children := getBaseChildren()
	iterations := 2

	var newParents Population = parents

	insertion := Add{}

	for i := 0; i < iterations; i++ {
		newParents = insertion.Insert(newParents, children)
	}

	assert.Equal(t, len(newParents), len(parents)+iterations*len(children))
}
