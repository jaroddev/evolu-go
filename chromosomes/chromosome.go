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
	child.Alleles = make([]bool, 0)

	for i := range chromosome.Alleles {
		child.Alleles = append(child.Alleles, chromosome.Alleles[i])
	}

	return child
}

func GetCopyOf(chromosome Chromosome) Chromosome {
	copy := GetChildOf(chromosome)

	copy.Age = chromosome.Age
	copy.Fitness = chromosome.Fitness

	return copy
}
