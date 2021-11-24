package insertions

import (
	. "evolugo/chromosomes"
	"sort"
)

func ReplaceIfBetter(pop Population, children Population) Population {
	sort.Slice(pop, func(i, j int) bool {
		return pop[i].Fitness < pop[j].Fitness
	})

	sort.Slice(children, func(i, j int) bool {
		return children[i].Fitness < children[j].Fitness
	})

	worsts := pop[0:len(children)]

	for worstIndex := range worsts {
		for index := 0; index < len(pop); index++ {
			// If the worst is the currently accessed population
			if &pop[index] == &worsts[worstIndex] {
				// And the worst's fitness is lower
				if children[worstIndex].Fitness > worsts[worstIndex].Fitness {
					// Replace it with the chldren
					pop[index] = children[worstIndex]
				}
				break
			}
		}
	}

	return pop
}
