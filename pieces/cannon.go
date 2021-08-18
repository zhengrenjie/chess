package pieces

type Cannon struct {
	PieceImpl
}

func (c *Cannon) Next() []*Point {
	ret := make([]*Point, 0)

	x0 := c.point.X()
	y0 := c.point.Y()

	carriage := false
	addPoint := func(cursor int) bool {
		piece, ok := c.Context().Metrics[cursor]
		if !carriage {
			if !ok {
				ret = append(ret, NewPoint(cursor, nil))
				return true
			} else {
				carriage = true
				return true
			}
		} else {
			if ok {
				if piece.Side() == -c.Side() {
					ret = append(ret, NewPoint(cursor, piece))
				}
				return false
			}
		}

		return true
	}

	/* go right */
	for x := x0; x <= Xmax; x++ {
		cursor := y0*10 + x
		if addPoint(cursor) {
			continue
		}
		break
	}

	/* go left */
	carriage = false
	for x := x0; x >= Xmin; x-- {
		cursor := y0*10 + x
		if addPoint(cursor) {
			continue
		}
		break
	}

	/* go up */
	carriage = false
	for y := y0; y <= Ymax; y++ {
		cursor := y*10 + x0
		if addPoint(cursor) {
			continue
		}
		break
	}

	/* go down */
	carriage = false
	for y := y0; y >= Ymin; y-- {
		cursor := y*10 + x0
		if addPoint(cursor) {
			continue
		}
		break
	}

	return ret
}
