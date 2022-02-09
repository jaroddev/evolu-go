package crossovers

import (
	"math/rand"

	. "github.com/jaroddev/evolugo/chromosomes"
)

type MonoPoint struct {
	ChildrenNumber int
}

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

func (c *MonoPoint) Cross(pop *Population) Population {
	children := make(Population, 0)

	for len(children) < c.ChildrenNumber {
		separationPoint := rand.Intn(len((*pop)[0].Alleles))

		firstParentIndex, secondParentIndex := chooseRandomlyTwoParentsFromAPopulation(len(*pop))
		firstParent, secondParent := (*pop)[firstParentIndex], (*pop)[secondParentIndex]

		firstChild := combineChromosomeFromIndex(firstParent, secondParent, separationPoint)
		secondChild := combineChromosomeFromIndex(secondParent, firstParent, separationPoint)

		children = append(children, firstChild, secondChild)
	}

	return children
}
