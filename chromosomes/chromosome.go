package chromosomes

type Chromosome struct {
	Age     int8
	Alleles []bool
	Fitness float64
}

type Population []Chromosome
