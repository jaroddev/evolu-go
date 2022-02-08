package genetic

import (
	. "evolugo/chromosomes"
	"math/rand"
)

type OneMax struct {
	InitPopLength int
	AlleleLength  int

	MaxCycle      int
	NotUpdatedFor int
}

// This is the base config to test the different configurations
// No need to configure this
func NewBasicConfig() *OneMax {
	return &OneMax{
		InitPopLength: 2,
		AlleleLength:  1000,
		MaxCycle:      1500,
		NotUpdatedFor: -1,
	}
}

func (problem *OneMax) Attach(algorithm *GA) {
	// problem specific
	algorithm.Init = problem.Init
	algorithm.Fit = problem.Fit
	algorithm.Continue = problem.Continue
}

func (o *OneMax) Init() Population {
	pop := make(Population, o.InitPopLength)

	for index := range pop {
		chromosome := NewChromosome()
		chromosome.Alleles = make([]bool, o.AlleleLength)
		pop[index] = chromosome
	}

	return pop
}

func (o *OneMax) RandomInit() Population {
	pop := make(Population, o.InitPopLength)

	for index := range pop {
		chromosome := NewChromosome()
		chromosome.Alleles = make([]bool, o.AlleleLength)

		for locus := range chromosome.Alleles {
			if rand.Intn(2) == 1 {
				chromosome.Alleles[locus] = true
			} else {
				chromosome.Alleles[locus] = false
			}
		}

		pop[index] = chromosome
	}

	return pop
}

func (o *OneMax) Fit(c *Chromosome) {
	c.Fitness = 0
	for _, allele := range c.Alleles {
		if allele {
			c.Fitness++
		}
	}
}

func (o *OneMax) Continue(algorithm *GA) bool {
	// Return true if the algorithm should continue
	// if return false then the algorithm stop

	// Stop if cycle number is higher than 80
	return algorithm.Generation < o.MaxCycle &&
		// or if there were no update for at least 20 cycles
		// algorithm.Generation-algorithm.LastUpdate < o.NotUpdatedFor &&
		// or if
		algorithm.Best.Fitness < float64(len(algorithm.Best.Alleles))
}
