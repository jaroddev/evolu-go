package mutations

import (
	. "evolugo/chromosomes"
	"fmt"
	"math/rand"
)

type Mutation func(*Chromosome)

func Flip(flips int) Mutation {

	return func(c *Chromosome) {
		if flips > len(c.Alleles) {
			panic("Not enough elements to flip")
		}

		indexes := make([]int, len(c.Alleles))
		for index := range indexes {
			indexes[index] = index
		}

		rand.Shuffle(len(indexes), func(i, j int) {
			indexes[i], indexes[j] = indexes[j], indexes[i]
		})

		picks := indexes[:flips]

		for _, pick := range picks {
			c.Alleles[pick] = !c.Alleles[pick]
		}
	}

}

// How to test this ??? like really ???
func BitFlip() Mutation {
	return func(c *Chromosome) {
		// If higher than this then flip the allele
		mutationProbability := 1 - 1/len(c.Alleles)

		fmt.Println("proba ", mutationProbability)

		// Try to mutate each flips
		for index := range c.Alleles {
			probabilityChromosomeMutate := rand.Float64()
			if probabilityChromosomeMutate > float64(mutationProbability) {
				c.Alleles[index] = !c.Alleles[index]
			}
		}
	}
}
