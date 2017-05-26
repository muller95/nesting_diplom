func (indiv *Individual) Mutate() (*Individual, error) {
	if len(indiv.Genom) < 2 {
		return nil, errors.New("Too short genom")
	}

	mutant := new(Individual)
	genomSize := len(indiv.Genom)
	mutant.Genom = make([]int, genomSize)
	copy(mutant.Genom, indiv.Genom)
	i := rand.Int() % genomSize
	j := rand.Int() % genomSize
	for i == j {
		j = rand.Int() % genomSize
	}
	mutant.Genom[i], mutant.Genom[j] = mutant.Genom[j], mutant.Genom[i]
	return mutant, nil
}