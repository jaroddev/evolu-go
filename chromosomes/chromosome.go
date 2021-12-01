package chromosomes

type Chromosome struct {
	Age     int
	Alleles []bool
	Fitness float64
}

func NewChromosome() Chromosome {
	return Chromosome{
		Age:     0,
		Alleles: nil,
		Fitness: 0,
	}
}

type Population []Chromosome
