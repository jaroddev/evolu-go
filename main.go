package main

import (
	. "evolugo/chromosomes"
	. "evolugo/genetic"
	. "evolugo/mutations"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	flipOption := func(algorithm *GA) *GA {
		algorithm.Mutation = Flip(1)

		return algorithm
	}

	algorithm := NewGA(flipOption)

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
		return g.Cycle < 200 &&
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
