package insertions

import (
	. "evolugo/chromosomes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getParents() Population {
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
		Chromosome{
			Age:     0,
			Fitness: 38,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
	}
}

func getChildren() Population {
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

func getOneGoodOneBadChildren() Population {
	return Population{
		Chromosome{
			Age:     0,
			Fitness: 150,
			Alleles: []bool{
				false, false, false, true, false, true, false, true,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: -2,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		}}
}

func getOneBadChild() Population {
	return Population{
		Chromosome{
			Age:     0,
			Fitness: -2,
			Alleles: []bool{
				true, false, false, true, false, true, false, true,
			},
		}}
}

func TestChildrenInserted(t *testing.T) {
	parents := getParents()
	children := getChildren()

	newParents := ReplaceIfBetter(parents, children)

	assert.Equal(t, len(parents), len(newParents))
	assert.Subset(t, newParents, children)
}

func TestOneReplaced(t *testing.T) {
	parents := getParents()
	children := getOneGoodOneBadChildren()
	childrenCopy := getOneGoodOneBadChildren()

	newParents := ReplaceIfBetter(parents, children)

	assert.Equal(t, len(parents), len(newParents))

	assert.Contains(t, newParents, childrenCopy[0])
	assert.NotContains(t, newParents, childrenCopy[1])
}

func TestNoneReplaced(t *testing.T) {
	parents := getParents()
	children := getOneBadChild()

	newParents := ReplaceIfBetter(parents, children)

	assert.Equal(t, len(parents), len(newParents))
	assert.NotContains(t, newParents, children[0])
}
