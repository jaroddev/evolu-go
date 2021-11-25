package mutations

import (
	. "evolugo/chromosomes"
	"testing"

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
	oneFlip := Flip(1)

	oneFlip(&chromosome)

	assert.Contains(t, chromosome.Alleles, true)

}

func TestThreeFlips(t *testing.T) {
	chromosome := getTestChromosome()

	flips := 3
	threeFlips := Flip(flips)
	threeFlips(&chromosome)

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
	tenFlips := Flip(flips)

	assert.Panics(t, func() {
		tenFlips(&chromosome)
	})
}

func TestBitFlipWithOneAlleledChromosome(t *testing.T) {
	oneAllele := Chromosome{
		Age:     0,
		Fitness: 0,
		Alleles: []bool{false},
	}

	flip := BitFlip()

	flip(&oneAllele)

	assert.Equal(t, len(oneAllele.Alleles), 1)
	assert.Contains(t, oneAllele.Alleles, true)
}
