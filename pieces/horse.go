package pieces

type Horse struct {
	Side  Side
	Point Point
}

func (h *Horse) Next() []*Point {
	ret := make([]*Point, 0)

	if h.Point.Y()-2 > 0 {
		_, ok := Metrics[h.Point.point-10]
		if !ok {
			addPoint(h.Point.point-11, -h.Side, &ret)
			addPoint(h.Point.point-9, -h.Side, &ret)
		}

	}

	return ret
}
