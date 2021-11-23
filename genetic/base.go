package genetic

import (
	. "evolugo/chromosomes"
	. "evolugo/crossovers"
	. "evolugo/mutations"
	. "evolugo/selections"

	"sort"
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
	Survive  func(*Population)
}

type GAOption func(*GA)

func NewGA(options ...GAOption) *GA {
	var algorithm GA

	for _, option := range options {
		option(&algorithm)
	}

	return &algorithm
}

func (algorithm *GA) Run() {
	// Always init cycle to zero
	algorithm.Cycle = 1

	algorithm.Pop = algorithm.Init()

	for index := range algorithm.Pop {
		algorithm.Fit(&algorithm.Pop[index])
	}

	algorithm.Best = &algorithm.Pop[0]

	for algorithm.Continue(algorithm) {
		// Check which is the new best chromosome

		parents := algorithm.Select(&algorithm.Pop)

		children := algorithm.Cross(&parents)

		for index := range children {
			algorithm.Mutation(&children[index])
			algorithm.Fit(&children[index])
		}

		algorithm.Pop = append(algorithm.Pop, children...)

		// descending sort
		sort.Slice(algorithm.Pop, func(i, j int) bool {
			return algorithm.Pop[j].Fitness < algorithm.Pop[i].Fitness
		})

		algorithm.Best = &algorithm.Pop[0]

		algorithm.Survive(&algorithm.Pop)

		algorithm.Cycle++
	}

}
