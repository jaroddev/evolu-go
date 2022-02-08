package main

import (
	. "evolugo/crossovers"
	"evolugo/genetic"
	. "evolugo/genetic"
	. "evolugo/insertions"
	. "evolugo/mutations"
	"evolugo/record"
	. "evolugo/selections"
	"math/rand"
	"time"
)

func main() {

	// prepare 30 seeds
	rand.Seed(time.Now().UnixNano())

	algorithm := &genetic.GA{}

	onemax := NewBasicConfig()
	onemax.Attach(algorithm)

	// parameters
	algorithm.Mutation = &Flip{Frequency: 1}
	algorithm.Selection = &SelectBest{ParentNumber: 2}
	algorithm.Crossover = &Clone{ChildrenNumber: 2}
	algorithm.Insertion = &Add{}

	recorder := record.NewRecorder()
	algorithm.Recorder = &recorder

	algorithm.Run()

	recorder.Save()

}
