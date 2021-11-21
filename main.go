package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Chromosome struct {
	Age     int8
	Alleles []bool
	Fitness float64
}

type Population []Chromosome

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

	Select   func(*Population) Population
	Cross    func(*Population) Population
	Mutation func(*Chromosome)
	Survive  func(*Population)
}

func (algorithm *GA) Run() {
	// Always init cycle to zero
	algorithm.Cycle = 1

	algorithm.Pop = algorithm.Init()

	for _, chromosome := range algorithm.Pop {
		algorithm.Fit(&chromosome)
	}

	algorithm.Best = &algorithm.Pop[0]

	for algorithm.Continue(algorithm) {
		// Check which is the new best chromosome

		parents := algorithm.Select(&algorithm.Pop)

		children := algorithm.Cross(&parents)

		for index := range children {
			algorithm.Fit(&algorithm.Pop[index])
		}

		algorithm.Pop = append(algorithm.Pop, children...)

		// descending sort
		sort.Slice(algorithm.Pop, func(i, j int) bool {
			return algorithm.Pop[j].Fitness < algorithm.Pop[i].Fitness
		})

		algorithm.Best = &algorithm.Pop[0]

		algorithm.Survive(&algorithm.Pop)

		algorithm.Cycle++

		fmt.Println("number of individuals in population", len(algorithm.Pop))
		fmt.Println("alleles of the best indiidual", algorithm.Best.Alleles)

		fmt.Println("best fitness", algorithm.Best.Fitness)
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())

	algorithm := &GA{}

	algorithm.Init = func() Population {
		initPopLength := 2
		alleleLength := 8

		pop := make(Population, initPopLength)

		for index := range pop {
			pop[index].Alleles = make([]bool, alleleLength)

			for locus := range pop[index].Alleles {
				if rand.Intn(2) == 1 {
					pop[index].Alleles[locus] = true
				} else {
					pop[index].Alleles[locus] = false
				}
			}
		}

		return pop
	}

	algorithm.Fit = func(c *Chromosome) {
		c.Fitness = 0
		for _, allele := range c.Alleles {
			if allele {
				c.Fitness++
			}
		}
	}

	algorithm.Continue = func(g *GA) bool {
		// Return true if the algorithm should continue
		// if return false then the algorithm stop

		// Stop if cycle number is higher than 80
		return g.Cycle < 80 &&
			// or if there were no update for at least 20 cycles
			g.Cycle-g.LastUpdate < 20 &&
			// or if
			g.Best.Fitness < float64(len(g.Best.Alleles))
	}

	algorithm.Select = func(p *Population) Population {
		parentNumber := 2

		if parentNumber > len((*p)) || parentNumber > cap((*p)) {
			return (*p)[:len((*p))]
		}

		return (*p)[:parentNumber]
	}

	algorithm.Cross = func(p *Population) Population {
		childrenNumber := 2

		if childrenNumber > len((*p)) || childrenNumber > cap((*p)) {
			return (*p)[:len((*p))]
		}

		return (*p)[:childrenNumber]
	}

	algorithm.Mutation = func(c *Chromosome) {
		flips := 3

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

	algorithm.Survive = func(p *Population) {
		fmt.Println("Survival method has yet to be implemented")
	}

	algorithm.Run()
}
