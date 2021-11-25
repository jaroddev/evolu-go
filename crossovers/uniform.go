package crossovers

import (
	. "evolugo/chromosomes"
	"math/rand"
)

func pickAParentWithUniformProbability(parents *Population) Chromosome {
	probability := rand.Float64()

	uniformProbability := 1 / float64(len(*parents))
	chosenParentIndex := int(probability / uniformProbability)

	return (*parents)[chosenParentIndex]
}

func Uniform(childrenNumber int) CrossOver {
	return func(pop *Population) Population {
		children := make(Population, 0)

		// Produce childrenNumber children
		for childIndex := 0; childIndex < childrenNumber; childIndex++ {

			// Create a child
			child := Chromosome{
				Age:     1,
				Fitness: 0,
				Alleles: (*pop)[childIndex].Alleles,
			}

			for alleleIndex := range child.Alleles {
				chosenParent := pickAParentWithUniformProbability(pop)
				child.Alleles[alleleIndex] = chosenParent.Alleles[alleleIndex]
			}

			children = append(children, child)
		}

		return children
	}
}
