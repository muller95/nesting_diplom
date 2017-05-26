func Crossover(parent1, parent2 *Individual) (*Individual, error) {
	genSize1 := len(parent1.Genom)
	genSize2 := len(parent2.Genom)
	if genSize1 != genSize2 {
		return nil, errors.New("Different sizes of genoms")
	}

	if genSize1 < 3 {
		return nil, errors.New("Too short genom")
	}

	g1 := rand.Int() % genSize1
	g2 := rand.Int() % genSize2

	child := new(Individual)
	child.Genom = make([]int, genSize1)
	child.Genom[g1] = parent1.Genom[g1]
	child.Genom[g2] = parent1.Genom[g2]

	for i, j := 0, 0; i < genSize2 && j < genSize2; i, j = i+1, j+1 {
		if j == g1 || j == g2 {
			i--
			continue
		}

		if parent2.Genom[i] == child.Genom[g1] || parent2.Genom[i] == child.Genom[g2] {
			j--
			continue
		}

		child.Genom[j] = parent2.Genom[i]
	}

	return child, nil
}
