package pieces

type Guard struct {
	PieceImpl
}

func (g *Guard) Next() []*Point {
	ret := make([]*Point, 0)
	point := g.Point().Point()

	if g.Side() == SideRed {
		switch point {
		case 3:
			addPoint(14, SideBlack, &ret)
		case 5:
			addPoint(14, SideBlack, &ret)
		case 14:
			addPoint(3, SideBlack, &ret)
			addPoint(5, SideBlack, &ret)
			addPoint(23, SideBlack, &ret)
			addPoint(25, SideBlack, &ret)
		case 23:
			addPoint(14, SideBlack, &ret)
		case 25:
			addPoint(14, SideBlack, &ret)
		}
	}

	if g.Side() == SideBlack {
		switch point {
		case 93:
			addPoint(84, SideRed, &ret)
		case 95:
			addPoint(84, SideRed, &ret)
		case 84:
			addPoint(93, SideRed, &ret)
			addPoint(95, SideRed, &ret)
			addPoint(73, SideRed, &ret)
			addPoint(75, SideRed, &ret)
		case 73:
			addPoint(84, SideRed, &ret)
		case 75:
			addPoint(84, SideRed, &ret)
		}
	}

	return ret
}
