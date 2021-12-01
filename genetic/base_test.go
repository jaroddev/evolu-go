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
	return NewGA(
		WithTemplates(&OneMax{
			InitPopLength: 2,
			AlleleLength:  8,
			MaxCycle:      20,
			NotUpdatedFor: 200,
		}),
	)
}

func getValidAlgorithm() *GA {
	algorithm := getValidOneMaxTemplate()

	algorithm.Select = SelectFirst(2)
	algorithm.Cross = CloneFirst(2)
	algorithm.Mutation = Flip(1)
	algorithm.Insert = insertions.ReplaceIfBetter

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
	algorithm.Select = nil

	assert.Panics(t, func() { algorithm.Run() })
}

func TestAlgorithmWithoutCross(t *testing.T) {
	algorithm := getValidAlgorithm()
	algorithm.Cross = nil

	assert.Panics(t, func() { algorithm.Run() })
}
func TestAlgorithmWithoutMutation(t *testing.T) {
	algorithm := getValidAlgorithm()
	algorithm.Mutation = nil

	assert.Panics(t, func() { algorithm.Run() })
}

func TestAlgorithmWithoutInsert(t *testing.T) {
	algorithm := getValidAlgorithm()
	algorithm.Insert = nil

	assert.Panics(t, func() { algorithm.Run() })
}
