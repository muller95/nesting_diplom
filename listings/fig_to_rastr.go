func (fig *Figure) FigToRastr(rt RastrType, resize int, bound int) *Rastr {
	rastr := new(Rastr)

	rastr.Width = int(fig.Width) + 1
	rastr.Height = int(fig.Height) + 1
	rastr.OuterContour = make([]PointInt, 0, rastr.Width*rastr.Height)
	rastr.RastrMatrix = make([][]int, rastr.Height)
	for i := 0; i < rastr.Height; i++ {
		rastr.RastrMatrix[i] = make([]int, rastr.Width)
	}

	for i := 0; i < len(fig.Primitives); i++ {
		for j := 0; j < len(fig.Primitives[i].Points)-1; j++ {

			var top, bottom Point

			if fig.Primitives[i].Points[j].Y > fig.Primitives[i].Points[j+1].Y {
				top = fig.Primitives[i].Points[j]
				bottom = fig.Primitives[i].Points[j+1]
			} else {
				top = fig.Primitives[i].Points[j+1]
				bottom = fig.Primitives[i].Points[j]
			}

			intervals := getIntervals(bottom.Y, top.Y)
			if top.Y-bottom.Y > 1.0 {
				for k := 0; k < len(intervals)-1; k++ {
					x1 := calcX(top, bottom, intervals[k])
					x2 := calcX(top, bottom, intervals[k+1])
					y := intervals[k]

					step := 1.0
					if x2 <= x1 {
						step = -1.0
					}
					rastr.RastrMatrix[int(y)][int(x1)] = filled
					rastr.RastrMatrix[int(y)][int(x2)] = filled
					for x := math.Trunc(x1); x != math.Trunc(x2); x += step {
						rastr.RastrMatrix[int(y)][int(x)] = filled
					}
				}
			} else {
				x1 := bottom.X
				x2 := top.X
				y := bottom.Y

				step := 1.0
				if x2 <= x1 {
					step = -1.0
				}
				rastr.RastrMatrix[int(y)][int(x1)] = filled
				rastr.RastrMatrix[int(y)][int(x2)] = filled
				for x := math.Trunc(x1); x != math.Trunc(x2); x += step {
					rastr.RastrMatrix[int(y)][int(x)] = filled
				}
			}
		}
	}

	if bound > 0 {
		rastr = makeBound(rastr, bound)
	}

	if resize > 0 {
		rastr = resizeRastr(rastr, resize)
	}

	rastr.findContour()
	for i := 0; i < rastr.Height; i++ {
		for j := 0; j < rastr.Width; j++ {
			if rastr.RastrMatrix[i][j] == contour {
				rastr.OuterContour = append(rastr.OuterContour, pointIntNew(j, i))
			}
		}
	}

	if rt == RastrTypePartInPart {
		rastr.floodRastrPartInPart()
	} else {
		rastr.floodRastrSimple()
	}

	return rastr
}