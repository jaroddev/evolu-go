package insertions

import (
	. "evolugo/chromosomes"
)

type Insertion func(parents Population, children Population) Population

func AddChildren(parents Population, children Population) Population {
	parents = append(parents, children...)
	return parents
}
