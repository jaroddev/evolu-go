package genetic

import (
	. "evolugo/chromosomes"
)

type Mutation interface {
	Mutate(c *Chromosome)
}

type Selection interface {
	Select(p *Population) Population
}

type CrossOver interface {
	Cross(p *Population) Population
}

type Insertion interface {
	Insert(Population, Population) Population
}

type GA struct {
	Pop Population
	// Best chromosome
	Best  *Chromosome
	Cycle int
	// Last cycle the
	LastUpdate int

	Init func() Population
	Fit  func(*Chromosome)

	Continue func(*GA) bool

	Selection Selection
	Crossover CrossOver
	Mutation  Mutation
	Insertion Insertion
}

func (algorithm *GA) cycleEnd() {
	algorithm.Cycle++
	for index := range algorithm.Pop {
		algorithm.Pop[index].Age++
	}
}

func (algorithm *GA) Run() {
	filter := &GetBest{}
	algorithm.Cycle = 0
	algorithm.Pop = algorithm.Init()

	for index := range algorithm.Pop {
		algorithm.Fit(&algorithm.Pop[index])
	}

	algorithm.Best = &FilterWithLimit(algorithm.Pop, filter, 1)[0]

	for algorithm.Continue(algorithm) {
		originalParent := algorithm.Selection.Select(&algorithm.Pop)
		parents := make(Population, len(originalParent))
		copy(parents, originalParent)

		children := algorithm.Crossover.Cross(&parents)

		for index := range children {
			algorithm.Mutation.Mutate(&children[index])
			algorithm.Fit(&children[index])
		}

		algorithm.Pop = algorithm.Insertion.Insert(algorithm.Pop, children)

		algorithm.Best = &FilterWithLimit(algorithm.Pop, filter, 1)[0]

		algorithm.cycleEnd()
	}

}
