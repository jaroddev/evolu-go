package main

import (
	. "evolugo/crossovers"
	. "evolugo/genetic"
	. "evolugo/insertions"
	. "evolugo/mutations"
	. "evolugo/selections"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	algorithm := NewGA(
		WithTemplates(&OneMax{
			InitPopLength: 2,
			AlleleLength:  8,
			MaxCycle:      100,
			NotUpdatedFor: -1,
		}),
	)

	algorithm.Mutation = &Flip{Frequency: 1}

	algorithm.Selection = &SelectFirst{ParentNumber: 2}

	// Too Similar to SelectFirst
	algorithm.Crossover = &Clone{ChildrenNumber: 2}

	algorithm.Insertion = &Add{}

	algorithm.Run()

	for index, chromosome := range algorithm.Pop {
		fmt.Println(index, " : ", chromosome)
	}

	fmt.Println("Best is => ", algorithm.Best)

}
