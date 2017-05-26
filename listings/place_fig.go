func placeFigHeight(fig *Figure, posits *[]Position, width, height, resize, bound int,
	place [][]int) bool {
	placed := false

	for angle := 0.0; angle < 360.0; angle += fig.AngleStep {

		currFig := fig.copy()
		currFig.Rotate(angle)
		rastr := currFig.figToRastr(rt, resize, bound)
		if rastr.Width > width/resize || rastr.Height > height/resize {
			return false
		}
		for y := 0; y < height-rastr.Height; y++ {
			for x := 0; x < width-rastr.Width; x++ {
				cross := false

				for k := 0; k < len(rastr.OuterContour); k++ {
					i, j := rastr.OuterContour[k].Y, rastr.OuterContour[k].X

					if place[y+i][x+j] > 0 {
						cross = true
						break
					}
				}

				if cross {
					continue
				}

				if checkPositionHeight(currFig, posits, float64(x*resize), float64(y*resize),
					float64(width), float64(height), &placed) {
					(*posits)[len(*posits)-1].Angle = angle
				}

				x = width
				y = height
			}
		}
	}

	if !placed {
		return false
	}

	rastr := (*posits)[len(*posits)-1].Fig.figToRastr(rt, resize, bound)
	for i := 0; i < rastr.Height; i++ {
		for j := 0; j < rastr.Width; j++ {
			x := int((*posits)[len(*posits)-1].X) / resize
			y := int((*posits)[len(*posits)-1].Y) / resize
			place[i+y][j+x] += rastr.RastrMatrix[i][j]
		}
	}

	return true
}