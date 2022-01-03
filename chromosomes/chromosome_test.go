package chromosomes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getSampleChromosome() Chromosome {
	return Chromosome{
		Age:     2,
		Fitness: 1,
		Alleles: []bool{true, false, true, false},
	}
}

func getSampledChild() Chromosome {
	return GetChildOf(getSampleChromosome())
}

func getSampledCopy() Chromosome {
	return GetCopyOf(getSampleChromosome())
}

func TestChildHasDefaultAge(t *testing.T) {
	child := getSampledChild()
	assert.Equal(t, child.Age, 0)
}

func TestChildHasDefaultFitness(t *testing.T) {
	child := getSampledChild()

	assert.Equal(t, child.Fitness, float64(0))
}

func TestChildHasSameAlleleValues(t *testing.T) {
	child := getSampledChild()

	assert.Equal(t, child.Alleles, []bool{true, false, true, false})
}

func TestCopyHasSameAge(t *testing.T) {
	sample := getSampleChromosome()
	copy := getSampledCopy()

	assert.Equal(t, sample.Age, copy.Age)
}

func TestCopyHasSameFitness(t *testing.T) {
	sample := getSampleChromosome()
	copy := getSampledCopy()

	assert.Equal(t, sample.Fitness, copy.Fitness)
}

func TestCopyHasSameAlleleValues(t *testing.T) {
	sample := getSampleChromosome()
	copy := getSampledCopy()

	assert.Equal(t, sample.Alleles, copy.Alleles)
}

func TestModifyingCopyWontAffectOriginal(t *testing.T) {
	sample := getSampleChromosome()
	copy := GetCopyOf(sample)

	copy.Alleles[0] = false

	assert.NotEqual(t, sample.Alleles, copy.Alleles)
}

func TestModifyingChildWontAffectOriginal(t *testing.T) {
	sample := getSampleChromosome()
	child := GetChildOf(sample)

	child.Alleles[0] = false

	assert.NotEqual(t, sample.Alleles, child.Alleles)
}
