package main

import (
	"sort"
)

type Chromosome struct {
	Age     int8
	Alleles []bool
	Fitness float64
}

type Population []Chromosome

type GA struct {
	Pop   Population
	Cycle int

	Init func() Population
	Fit  func(*Chromosome)

	Stop func(*Population, int) bool

	Select   func(*Population) Population
	Cross    func(*Population) Population
	Mutation func(*Chromosome)
	Survive  func(*Population)
}

func (algorithm *GA) Run() {

	algorithm.Pop = algorithm.Init()

	for _, chromosome := range algorithm.Pop {
		algorithm.Fit(&chromosome)
	}

	for algorithm.Stop(&algorithm.Pop, algorithm.Cycle) {

		parents := algorithm.Select(&algorithm.Pop)

		children := algorithm.Cross(&parents)

		for _, child := range children {
			algorithm.Fit(&child)
		}

		algorithm.Pop = append(algorithm.Pop, children...)

		// descending sort
		sort.Slice(algorithm.Pop, func(i, j int) bool {
			return algorithm.Pop[j].Fitness < algorithm.Pop[i].Fitness
		})

	}

}
