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

func GetChildOf(chromosome Chromosome) Chromosome {
	child := NewChromosome()
	child.Fitness = chromosome.Fitness

	child.Alleles = make([]bool, len(chromosome.Alleles))
	copy(child.Alleles, chromosome.Alleles)

	return child
}

func GetCopyOf(chromosome Chromosome) Chromosome {
	copy := GetChildOf(chromosome)

	copy.Age = chromosome.Age
	copy.Fitness = chromosome.Fitness

	return copy
}
