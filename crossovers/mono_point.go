package crossovers

import (
	. "evolugo/chromosomes"
	"math/rand"
)

func combineChromosomeFromIndex(firstParent Chromosome, secondParent Chromosome, separationPoint int) Chromosome {
	child := NewChromosome()
	child.Alleles = make([]bool, 0)

	for i := 0; i < separationPoint; i++ {
		child.Alleles = append(child.Alleles, firstParent.Alleles[i])
	}

	for i := separationPoint; i < len(secondParent.Alleles); i++ {
		child.Alleles = append(child.Alleles, secondParent.Alleles[i])
	}

	return child
}

func chooseRandomlyTwoParentsFromAPopulation(parentNumber int) (int, int) {
	randomIndexPermutation := rand.Perm(parentNumber)
	firstParentIndex, secondParentIndex := randomIndexPermutation[0], randomIndexPermutation[1]

	return firstParentIndex, secondParentIndex
}

func MonoPoint(childrenNumber int) CrossOver {
	return func(parents *Population) Population {
		children := make(Population, 0)

		for len(children) < childrenNumber {
			separationPoint := rand.Intn(len((*parents)[0].Alleles))

			firstParentIndex, secondParentIndex := chooseRandomlyTwoParentsFromAPopulation(len(*parents))
			firstParent, secondParent := (*parents)[firstParentIndex], (*parents)[secondParentIndex]

			firstChild := combineChromosomeFromIndex(firstParent, secondParent, separationPoint)
			secondChild := combineChromosomeFromIndex(secondParent, firstParent, separationPoint)

			children = append(children, firstChild, secondChild)
		}

		return children
	}
}
