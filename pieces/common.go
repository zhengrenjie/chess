package pieces

func addPoint(point int, side Side, points *[]*Point) {
	if piece, ok := Metrics[point]; ok {
		if piece.Side() == side {
			*points = append(*points, &Point{point: point, piece: piece})
		}
	} else {
		*points = append(*points, &Point{point: point})
	}
}
