package pieces

type Cannon struct {
	Side  Side
	Point Point
}

func (c *Cannon) Next() []*Point {
	ret := make([]*Point, 0)

	x0 := c.Point.x
	y0 := c.Point.y

	carriage := false
	addPoint := func(cursor int) bool {
		piece, ok := Metrics[cursor]
		if !carriage {
			if ok {
				ret = append(ret, NewPoint(cursor, nil))
				return true
			} else {
				carriage = true
				return true
			}
		} else {
			if ok {
				if piece.Side() == -c.Side {
					ret = append(ret, NewPoint(cursor, piece))
				}
				return false
			}
		}

		return true
	}

	/* go right */
	for x := c.Point.X(); x <= Xmax; x++ {
		cursor := y0*10 + x
		if addPoint(cursor) {
			continue
		}
		break
	}

	/* go left */
	carriage = false
	for x := c.Point.X(); x >= Xmin; x-- {
		cursor := y0*10 + x
		if addPoint(cursor) {
			continue
		}
		break
	}

	/* go up */
	carriage = false
	for y := c.Point.Y(); y <= Ymax; y++ {
		cursor := y*10 + x0
		if addPoint(cursor) {
			continue
		}
		break
	}

	/* go down */
	carriage = false
	for y := c.Point.Y(); y <= Ymax; y++ {
		cursor := y*10 + x0
		if addPoint(cursor) {
			continue
		}
		break
	}

	return ret
}
