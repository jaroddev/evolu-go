package main

import (
	. "evolugo/chromosomes"
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

	algorithm.Mutation = Flip(1)

	algorithm.Select = SelectFirst(2)

	// Too Similar to SelectFirst
	algorithm.Cross = CloneFirst(2)

	algorithm.Insert = AddChildren

	algorithm.Survive = func(p *Population) {}

	algorithm.Run()

	for index, chromosome := range algorithm.Pop {
		fmt.Println(index, " : ", chromosome)
	}
}
