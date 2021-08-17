package pieces

const (
	Xmin = 0
	Xmax = 8
	Ymin = 0
	Ymax = 9
)

func addPoint(point int, side Side, points *[]*Point) {
	if piece, ok := Metrics[point]; ok {
		if piece.Side() == side {
			*points = append(*points, NewPoint(point, piece))
		}
	} else {
		*points = append(*points, NewPoint(point, nil))
	}
}
