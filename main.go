package main

import (
	. "evolugo/chromosomes"
	. "evolugo/crossovers"
	. "evolugo/genetic"
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
			MaxCycle:      200,
			NotUpdatedFor: 200,
		}),
	)

	algorithm.Mutation = Flip(10)

	algorithm.Select = SelectFirst(2)

	// Too Similar to SelectFirst
	algorithm.Cross = CloneFirst(2)

	algorithm.Survive = func(p *Population) {}

	algorithm.Run()

	for index, chromosome := range algorithm.Pop {
		fmt.Println(index, " : ", chromosome)
	}
}
