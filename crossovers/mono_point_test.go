package crossovers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonoPointOnTwoChromosomes(t *testing.T) {
	parents := getUniformParent()

	firstParent, secondParent := parents[0], parents[2]

	separationPoint := 2

	child := combineChromosomeFromIndex(firstParent, secondParent, separationPoint)

	assert.Equal(t, firstParent.Alleles[:separationPoint], child.Alleles[:separationPoint])
	assert.NotEqual(t, secondParent.Alleles[:separationPoint], child.Alleles[:separationPoint])

	assert.NotEqual(t, firstParent.Alleles[separationPoint:], child.Alleles[separationPoint:])
	assert.Equal(t, secondParent.Alleles[separationPoint:], child.Alleles[separationPoint:])
}
