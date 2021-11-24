package insertions

import (
	. "evolugo/chromosomes"
)

func getOldest(pop Population, numberToReplace int) []int {
	indexesToBeReplaced := make([]int, 0)

	for len(indexesToBeReplaced) < numberToReplace {
		oldest := -1
		for index, chromosome := range pop {
			alreadyMarkedAsToBeReplaced := false

			for _, indexToBeReplaced := range indexesToBeReplaced {
				if index == indexToBeReplaced {
					alreadyMarkedAsToBeReplaced = true
					break
				}
			}

			if alreadyMarkedAsToBeReplaced {
				continue
			}

			if oldest == -1 {
				oldest = index
				continue
			}

			if chromosome.Age > pop[oldest].Age {
				oldest = index
			}
		}
		indexesToBeReplaced = append(indexesToBeReplaced, oldest)
	}

	return indexesToBeReplaced
}

func ReplaceOldest(pop Population, children Population) Population {

	indexesToBeReplaced := getOldest(pop, len(children))

	for index := 0; index < len(indexesToBeReplaced); index++ {
		indexToBeReplaced := indexesToBeReplaced[index]
		pop[indexToBeReplaced] = children[index]
	}

	return pop
}
