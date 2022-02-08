package genetic

import (
	. "evolugo/crossovers"
	"evolugo/insertions"
	. "evolugo/mutations"
	. "evolugo/selections"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getValidOneMaxTemplate() *GA {
	onemax := &OneMax{
		InitPopLength: 2,
		AlleleLength:  8,
		MaxCycle:      20,
		NotUpdatedFor: 200,
	}

	algorithm := &GA{}
	onemax.Attach(algorithm)

	return algorithm
}

func getValidAlgorithm() *GA {
	algorithm := getValidOneMaxTemplate()

	algorithm.Selection = &SelectBest{
		ParentNumber: 2,
	}
	algorithm.Crossover = &Clone{ChildrenNumber: 2}
	algorithm.Mutation = &Flip{
		Frequency: 1,
	}
	algorithm.Insertion = &insertions.Elitist{}

	return algorithm
}

func TestDefaultAlgorithmPanic(t *testing.T) {
	algorithm := GA{}
	assert.Panics(t, func() { algorithm.Run() })
}

func TestAlgorithmWithTemplatePanic(t *testing.T) {
	algorithm := getValidOneMaxTemplate()

	assert.Panics(t, func() { algorithm.Run() })
}

func TestAlgorithmWithoutSelection(t *testing.T) {
	algorithm := getValidAlgorithm()
	algorithm.Selection = nil

	assert.Panics(t, func() { algorithm.Run() })
}

func TestAlgorithmWithoutCross(t *testing.T) {
	algorithm := getValidAlgorithm()
	algorithm.Crossover = nil

	assert.Panics(t, func() { algorithm.Run() })
}
func TestAlgorithmWithoutMutation(t *testing.T) {
	algorithm := getValidAlgorithm()
	algorithm.Mutation = nil

	assert.Panics(t, func() { algorithm.Run() })
}

func TestAlgorithmWithoutInsert(t *testing.T) {
	algorithm := getValidAlgorithm()
	algorithm.Insertion = nil

	assert.Panics(t, func() { algorithm.Run() })
}
