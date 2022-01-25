package mutations

import (
	. "evolugo/chromosomes"
	"math/rand"
)

// Most basic mutationn needs to be parametered
type Flip struct {
	// Number of flip done for each mutation
	Frequency int
}

func (mutation *Flip) Mutate(c *Chromosome) {

	if mutation.Frequency > len(c.Alleles) {
		panic("Not enough elements to flip")
	}

	indexes := rand.Perm(len(c.Alleles))
	picks := indexes[:mutation.Frequency]

	for _, pick := range picks {
		c.Alleles[pick] = !c.Alleles[pick]
	}

}

type BitFlip struct{}

func (mutation *BitFlip) Mutate(c *Chromosome) {

	// If higher than this then flip the allele
	mutationProbability := 1 - 1/len(c.Alleles)

	// Try to mutate each flips
	for index := range c.Alleles {
		probabilityChromosomeMutate := rand.Float64()
		if probabilityChromosomeMutate > float64(mutationProbability) {
			c.Alleles[index] = !c.Alleles[index]
		}
	}

}
