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
		_, ok := h.Context().Metrics[point-10]
		if !ok {
			if x0 > Xmin {
				addPoint(point-20-1, -h.Side(), h.Context(), &ret)
			}
			if x0 < Xmax {
				addPoint(point-20+1, -h.Side(), h.Context(), &ret)
			}
		}
	}

	/* jump up */
	if y0+2 <= Ymax {
		_, ok := h.Context().Metrics[point+10]
		if !ok {
			if x0 > Xmin {
				addPoint(point+20-1, -h.Side(), h.Context(), &ret)
			}
			if x0 < Xmax {
				addPoint(point+20+1, -h.Side(), h.Context(), &ret)
			}
		}
	}

	/* jump left */
	if x0-2 >= Xmin {
		_, ok := h.Context().Metrics[point-1]
		if !ok {
			if y0 > Ymin {
				addPoint(point-2-10, -h.Side(), h.Context(), &ret)
			}
			if y0 < Ymax {
				addPoint(point-2+10, -h.Side(), h.Context(), &ret)
			}
		}
	}

	/* jump right */
	if x0+2 <= Xmax {
		_, ok := h.Context().Metrics[point+1]
		if !ok {
			if y0 > Ymin {
				addPoint(point+2-10, -h.Side(), h.Context(), &ret)
			}
			if y0 < Ymax {
				addPoint(point+2+10, -h.Side(), h.Context(), &ret)
			}
		}
	}

	return ret
}
