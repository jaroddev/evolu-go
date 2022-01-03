package chromosomes

type Filter interface {
	// Two Chromosome: candidate and currentSelection
	// function check if candidate is better than currentSelection
	IsBetterFit(*Chromosome, *Chromosome) bool
}

type GetBest struct{}

func (rule *GetBest) IsBetterFit(candidate *Chromosome, currentlySelected *Chromosome) bool {
	return candidate.Fitness > currentlySelected.Fitness
}

type GetOldest struct{}

func (rule *GetOldest) IsBetterFit(candidate *Chromosome, currentlySelected *Chromosome) bool {
	return candidate.Age > currentlySelected.Age
}
