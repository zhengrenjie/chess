package pieces

type Horse struct {
	Side  Side
	Point Point
}

func (h *Horse) Next() []*Point {
	ret := make([]*Point, 0)

	/* jump down */
	if h.Point.Y()-2 >= Ymin {
		_, ok := Metrics[h.Point.point-10]
		if !ok {
			if h.Point.X() > Xmin {
				addPoint(h.Point.point-20-1, -h.Side, &ret)
			}
			if h.Point.X() < Xmax {
				addPoint(h.Point.point-20+1, -h.Side, &ret)
			}
		}
	}

	/* jump up */
	if h.Point.Y()+2 <= Ymax {
		_, ok := Metrics[h.Point.point+10]
		if !ok {
			if h.Point.X() > Xmin {
				addPoint(h.Point.point+20-1, -h.Side, &ret)
			}
			if h.Point.X() < Xmax {
				addPoint(h.Point.point+20+1, -h.Side, &ret)
			}
		}
	}

	/* jump left */
	if h.Point.X()-2 >= Xmin {
		_, ok := Metrics[h.Point.point-1]
		if !ok {
			if h.Point.Y() > Ymin {
				addPoint(h.Point.point-2-10, -h.Side, &ret)
			}
			if h.Point.Y() < Ymax {
				addPoint(h.Point.point-2+10, -h.Side, &ret)
			}
		}
	}

	/* jump right */
	if h.Point.X()+2 <= Xmax {
		_, ok := Metrics[h.Point.point+1]
		if !ok {
			if h.Point.Y() > Ymin {
				addPoint(h.Point.point+2-10, -h.Side, &ret)
			}
			if h.Point.Y() < Ymax {
				addPoint(h.Point.point+2+10, -h.Side, &ret)
			}
		}
	}

	return ret
}
