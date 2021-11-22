package mutations

import (
	. "evolugo/chromosomes"
	"math/rand"
)

type Mutation func(*Chromosome)

func Flip(flips int) Mutation {

	return func(c *Chromosome) {
		if len(c.Alleles) <= flips {
			panic("Not enough elements to flip")
		}

		indexes := make([]int, len(c.Alleles))
		for index := range indexes {
			indexes[index] = index
		}

		rand.Shuffle(len(indexes), func(i, j int) {
			indexes[i], indexes[j] = indexes[j], indexes[i]
		})

		picks := indexes[:flips+1]

		for _, pick := range picks {
			c.Alleles[pick] = !c.Alleles[pick]
		}
	}

}
