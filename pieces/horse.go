package pieces

type Horse struct {
	PieceImpl
}

func (h *Horse) Next() []*Point {
	ret := make([]*Point, 0)

	point := h.Point().Point()
	x0 := h.Point().X()
	y0 := h.Point().Y()

	/* jump down */
	if y0-2 >= Ymin {
		_, ok := Metrics[point-10]
		if !ok {
			if x0 > Xmin {
				addPoint(point-20-1, -h.Side(), &ret)
			}
			if x0 < Xmax {
				addPoint(point-20+1, -h.Side(), &ret)
			}
		}
	}

	/* jump up */
	if y0+2 <= Ymax {
		_, ok := Metrics[point+10]
		if !ok {
			if x0 > Xmin {
				addPoint(point+20-1, -h.Side(), &ret)
			}
			if x0 < Xmax {
				addPoint(point+20+1, -h.Side(), &ret)
			}
		}
	}

	/* jump left */
	if x0-2 >= Xmin {
		_, ok := Metrics[point-1]
		if !ok {
			if y0 > Ymin {
				addPoint(point-2-10, -h.Side(), &ret)
			}
			if y0 < Ymax {
				addPoint(point-2+10, -h.Side(), &ret)
			}
		}
	}

	/* jump right */
	if x0+2 <= Xmax {
		_, ok := Metrics[point+1]
		if !ok {
			if y0 > Ymin {
				addPoint(point+2-10, -h.Side(), &ret)
			}
			if y0 < Ymax {
				addPoint(point+2+10, -h.Side(), &ret)
			}
		}
	}

	return ret
}
