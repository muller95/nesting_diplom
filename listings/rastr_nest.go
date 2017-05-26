func RastrNest(figSet []*Figure, indiv *Individual, width, height, bound, resize int,
	rastrType RastrType, placementMode PlacementMode) error {
	if width <= 0 {
		return errors.New("Negative or zero width")
	} else if height <= 0 {
		return errors.New("Negative or zero height")
	} else if resize < 0 {
		return errors.New("Negative resize")
	} else if bound < 0 {
		return errors.New("Negative bound")
	}

	if resize < 1 {
		resize = 1
	}

	posits := make([]Position, 0)
	place := make([][]int, height/resize)
	for i := 0; i < height/resize; i++ {
		place[i] = make([]int, width/resize)
	}

	if len(indiv.Genom) == 0 {
		indiv.Genom = make([]int, 0)
	}

	mask := make([]int, len(figSet))
	failNest := make(map[int]bool)
	for i := 0; i < len(indiv.Genom); i++ {
		figNum := indiv.Genom[i]
		fig := figSet[figNum]
		if failNest[fig.ID] {
			continue
		}
		if placeFigHeight(fig, &posits, width, height, resize,
			bound, place) {
			posits[len(posits)-1].Fig.Translate(posits[len(posits)-1].X, posits[len(posits)-1].Y)
			mask[i] = 1
		} else {
			// fmt.Println("Fail nest")
			failNest[fig.ID] = true
		}
	}

	if len(posits) < len(indiv.Genom) {
		indiv.Height = math.Inf(1)
		return nil
	}

	for i := 0; i < len(figSet); i++ {
		fig := figSet[i]
		if mask[i] > 0 || failNest[fig.ID] {
			continue
		}
		if placeFigHeight(fig, &posits, width, height, resize,
			bound, place) {
			posits[len(posits)-1].Fig.Translate(posits[len(posits)-1].X, posits[len(posits)-1].Y)
			indiv.Genom = append(indiv.Genom, i)
		} else {
			// fmt.Println("Fail nest")
			failNest[fig.ID] = true
		}
	}

	indiv.Positions = posits
	maxHeight := 0.0
	for i := 0; i < len(posits); i++ {
		currHeight := posits[i].X + posits[i].Fig.Height
		maxHeight = math.Max(currHeight, maxHeight)
	}
	indiv.Height = maxHeight
	return nil
}