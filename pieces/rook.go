package pieces

type Rook struct {
	PieceImpl
}

func (r *Rook) Type() PieceType {
	return TypeRook
}

func (r *Rook) String() string {
	return "車"
}

func (r *Rook) Next() []*Point {
	ret := make([]*Point, 0)

	x0 := r.Point().X()
	y0 := r.Point().Y()

	addPoint := func(cursor int) bool {
		piece, ok := r.Context().Metrics[cursor]
		if ok {
			if piece.Side() == -r.Side() {
				ret = append(ret, NewPoint(cursor, piece))
			}

			return false
		} else {
			ret = append(ret, NewPoint(cursor, nil))
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
	for x := x0; x >= Xmin; x-- {
		cursor := y0*10 + x
		if addPoint(cursor) {
			continue
		}
		break
	}

	/* go up */
	for y := y0; y <= Ymax; y++ {
		cursor := y*10 + x0
		if addPoint(cursor) {
			continue
		}
		break
	}

	/* go down */
	for y := y0; y >= Ymin; y-- {
		cursor := y*10 + x0
		if addPoint(cursor) {
			continue
		}
		break
	}

	return ret
}
