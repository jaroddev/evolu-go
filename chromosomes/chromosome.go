package chromosomes

type Chromosome struct {
	Age     int
	Alleles []bool
	Fitness float64
}

type Population []Chromosome
