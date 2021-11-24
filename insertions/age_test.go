package insertions

import (
	. "evolugo/chromosomes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getOldParents() Population {
	return Population{
		Chromosome{
			Age:     25,
			Fitness: 10,
			Alleles: []bool{
				true, false, false, true, false, true, false, false,
			},
		},
		Chromosome{
			Age:     250,
			Fitness: 150,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     1200,
			Fitness: -2,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     1,
			Fitness: -2,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     6,
			Fitness: 1,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     78,
			Fitness: 38,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
	}
}

func getNewChildren() Population {
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
		}}
}

func TestGetThreeOldest(t *testing.T) {
	pop := getOldParents()
	indexes := getOldest(pop, 3)

	assert.Equal(t, len(indexes), 3)
	assert.Contains(t, indexes, 1)
	assert.Contains(t, indexes, 5)
	assert.Contains(t, indexes, 2)
}

func TestGetTwoOldest(t *testing.T) {
	pop := getOldParents()
	indexes := getOldest(pop, 2)

	assert.Equal(t, len(indexes), 2)
	assert.Contains(t, indexes, 1)
	assert.Contains(t, indexes, 2)
}

func TestGetOneOldest(t *testing.T) {
	pop := getOldParents()
	indexes := getOldest(pop, 1)

	assert.Equal(t, len(indexes), 1)
	assert.Contains(t, indexes, 2)
}

func TestOldestReplaced(t *testing.T) {
	parents := getOldParents()
	children := getNewChildren()

	newParents := ReplaceOldest(parents, children)

	assert.Equal(t, len(parents), len(newParents))
	assert.Subset(t, newParents, children)
}
