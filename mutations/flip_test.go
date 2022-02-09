package mutations

import (
	"testing"

	. "github.com/jaroddev/evolugo/chromosomes"

	"github.com/stretchr/testify/assert"
)

func getTestChromosome() Chromosome {
	return Chromosome{
		Age:     0,
		Fitness: 0.0,
		Alleles: []bool{
			false, false, false, false, false, false, false, false,
		},
	}
}

func TestDefaultTestChromosomeHaveNoTrue(t *testing.T) {
	chromosome := getTestChromosome()
	assert.NotContains(t, chromosome.Alleles, true)
}

func TestOneFlip(t *testing.T) {
	chromosome := getTestChromosome()
	oneFlip := &Flip{
		Frequency: 1,
	}

	oneFlip.Mutate(&chromosome)

	assert.Contains(t, chromosome.Alleles, true)

}

func TestThreeFlips(t *testing.T) {
	chromosome := getTestChromosome()

	flips := 3
	threeFlips := &Flip{
		Frequency: flips,
	}
	threeFlips.Mutate(&chromosome)

	fitness := 0
	for _, allele := range chromosome.Alleles {
		if allele {
			fitness++
		}
	}

	assert.Contains(t, chromosome.Alleles, true)
	assert.Equal(t, fitness, flips)
}

func TestPanicWhenMoreFlipsThanAlleles(t *testing.T) {
	chromosome := getTestChromosome()

	flips := 10
	tenFlips := &Flip{
		Frequency: flips,
	}

	assert.Panics(t, func() {
		tenFlips.Mutate(&chromosome)
	})
}

func TestBitFlipWithOneAlleledChromosome(t *testing.T) {
	oneAllele := Chromosome{
		Age:     0,
		Fitness: 0,
		Alleles: []bool{false},
	}

	flip := &BitFlip{}

	flip.Mutate(&oneAllele)

	assert.Equal(t, len(oneAllele.Alleles), 1)
	assert.Contains(t, oneAllele.Alleles, true)
}
