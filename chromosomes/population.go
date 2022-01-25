package chromosomes

import (
	"encoding/json"
	"os"
)

type Population []Chromosome

func NewWithFilename(filename string) (Population, error) {

	pop := make(Population, 0)

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &pop)

	if err != nil {
		return nil, err
	}

	return pop, nil

}

func insertAt(chromosome Chromosome, pop Population, index int) Population {

	firstPosition, lastPosition := 0, len(pop)

	if index == firstPosition {
		return append(Population{chromosome}, pop...)
	}

	if index == lastPosition {
		return append(pop, chromosome)
	}

	pop = append(pop[:index],
		append(Population{chromosome}, pop[index:]...)...)

	return pop
}

func FilterWithLimit(pop Population, selector Filter, numberOfSelection int) Population {

	if numberOfSelection > len(pop) {
		panic("Too many selection were asked")
	}

	selections := append(
		Population{},
		GetCopyOf(pop[0]),
	)

	for _, chromosome := range pop[1:] {

		selected := 0
		worst := true

		// Insertion
		for ; selected < len(selections); selected++ {

			if selector.IsBetterFit(&chromosome, &selections[selected]) {
				worst = false
				break
			}

		}

		newElement := GetCopyOf(chromosome)

		if worst {
			selected = len(selections)
		}

		selections = insertAt(
			newElement,
			selections,
			selected,
		)

	}

	return selections[:numberOfSelection]
}
