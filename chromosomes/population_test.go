package chromosomes

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePopulationWithFileThatDoNotExist(t *testing.T) {
	filename := "../fixtures/donotexist.json"
	pop, err := NewWithFilename(filename)

	assert.NoFileExists(t, filename)
	assert.NotNil(t, err)
	assert.Nil(t, pop)

}

func TestCreatePopulationWithFile(t *testing.T) {
	filename := "../fixtures/single_chromosome.json"
	pop, err := NewWithFilename(filename)

	assert.FileExists(t, filename)
	assert.Nil(t, err)

	assert.NotNil(t, pop)
	assert.Equal(t, len(pop) > 0, true)
}

func TestInsertionAtIndex(t *testing.T) {

	getPop := func() Population {
		return Population{
			{
				Age:     1,
				Fitness: 25,
			},
			{
				Age:     1,
				Fitness: 50,
			},
			{
				Age:     1,
				Fitness: 51,
			},
			{
				Age:     1,
				Fitness: 39,
			},
		}
	}

	pop := getPop()

	for position := range rand.Perm(len(pop) + 1) {

		t.Run(fmt.Sprintf("insertion at %d", position), func(t *testing.T) {

			chromosome := NewChromosome()
			chromosome.Age = 4
			chromosome.Fitness = 9.0

			newPop := insertAt(chromosome, pop, position)

			assert.Len(t, newPop, 5)
			assert.Equal(t, newPop[position], chromosome)

			for c := range pop {
				pop[c].Fitness += 30
			}

			fmt.Println("pop is ", pop, " and new Pop is ", newPop)

			// reset pop
			pop = getPop()
		})

	}
}

func TestBestThreeElements(t *testing.T) {

	getPop := func() Population {
		return Population{
			{
				Age:     1,
				Fitness: 25,
			},
			{
				Age:     1,
				Fitness: 50,
			},
			{
				Age:     1,
				Fitness: 51,
			},
			{
				Age:     1,
				Fitness: 39,
			},
			{
				Age:     3,
				Fitness: 8,
			},
			{
				Age:     3,
				Fitness: 10,
			},
			{
				Age:     6,
				Fitness: 451,
			},
			{
				Age:     6,
				Fitness: 2901,
			},
		}
	}

	selectionExpectedLength := 3

	pop := getPop()

	selection := FilterWithLimit(pop, &GetBest{}, selectionExpectedLength)

	assert.Len(t, selection, 3)

	assert.Equal(t, selection[0].Fitness, 2901.0)
	assert.Equal(t, selection[1].Fitness, 451.0)
	assert.Equal(t, selection[2].Fitness, 51.0)

}
