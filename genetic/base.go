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

type Recorder interface {
	Record(Generation int, BestFitness float64)
}

type GA struct {
	Pop Population
	// Best chromosome
	Best       *Chromosome
	Generation int
	// Last cycle the
	LastUpdate int

	Init func() Population
	Fit  func(*Chromosome)

	Continue func(*GA) bool

	Selection Selection
	Crossover CrossOver
	Mutation  Mutation
	Insertion Insertion

	// Remember bestfitness for each generation
	Recorder Recorder
}

func (algorithm *GA) updateBest() {
	updated := false

	if algorithm.Best == nil {
		algorithm.Best = &algorithm.Pop[0]
	}

	for index := range algorithm.Pop {
		if algorithm.Best.Fitness < algorithm.Pop[index].Fitness {
			algorithm.Best = &algorithm.Pop[index]
			updated = true
		}
	}

	if updated {
		algorithm.LastUpdate = 0
	} else {
		algorithm.LastUpdate++
	}
}

func (algorithm *GA) cycleEnd() {
	algorithm.Generation++
	for index := range algorithm.Pop {
		algorithm.Pop[index].Age++
	}
}

func (algorithm *GA) newGeneration() {
	originalParent := algorithm.Selection.Select(&algorithm.Pop)
	parents := make(Population, 0)

	// reset age of the parents
	for index := range originalParent {
		parents = append(
			parents,
			GetChildOf(originalParent[index]),
		)
	}

	children := algorithm.Crossover.Cross(&parents)

	for index := range children {
		algorithm.Mutation.Mutate(&children[index])
		algorithm.Fit(&children[index])
	}

	algorithm.Pop = algorithm.Insertion.Insert(algorithm.Pop, children)
}

func (algorithm *GA) Run() {
	algorithm.Pop = algorithm.Init()

	for index := range algorithm.Pop {
		algorithm.Fit(&algorithm.Pop[index])
	}

	algorithm.updateBest()

	for algorithm.Continue(algorithm) {
		algorithm.newGeneration()
		algorithm.updateBest()
		algorithm.cycleEnd()
		algorithm.Recorder.Record(algorithm.Generation, algorithm.Best.Fitness)
	}
}
