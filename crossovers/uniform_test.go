package crossovers

import (
	. "evolugo/chromosomes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getUniformParent() Population {

	return Population{
		Chromosome{
			Age:     0,
			Fitness: 0,
			Alleles: []bool{
				true, true, false, false, false, false,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 0,
			Alleles: []bool{
				false, false, true, true, false, false,
			},
		},
		Chromosome{
			Age:     0,
			Fitness: 0,
			Alleles: []bool{
				false, false, false, false, true, true,
			},
		},
	}
}

func TestUniform(t *testing.T) {
	parents := getUniformParent()
	uniformMutation := Uniform(2)

	children := uniformMutation(&parents)

	assert.Equal(t, len(children), 2)
}
