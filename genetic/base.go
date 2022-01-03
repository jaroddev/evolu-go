package genetic

import (
	. "evolugo/chromosomes"
	. "evolugo/crossovers"
	. "evolugo/insertions"
	. "evolugo/mutations"
	. "evolugo/selections"
)

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

	Select   Selection
	Cross    CrossOver
	Mutation Mutation
	Insert   Insertion
}

type GAOption func(*GA)

func NewGA(options ...GAOption) *GA {
	var algorithm GA

	for _, option := range options {
		option(&algorithm)
	}

	return &algorithm
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
		parents := algorithm.Select(&algorithm.Pop)

		copy(parents, parents)

		children := algorithm.Cross(&parents)

		for index := range children {
			algorithm.Mutation(&children[index])
			algorithm.Fit(&children[index])
		}

		algorithm.Pop = algorithm.Insert(algorithm.Pop, children)

		algorithm.Best = &FilterWithLimit(algorithm.Pop, filter, 1)[0]

		algorithm.cycleEnd()
	}

}
