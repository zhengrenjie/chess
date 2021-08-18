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
			addPoint(14, SideBlack, g.Context(), &ret)
		case 5:
			addPoint(14, SideBlack, g.Context(), &ret)
		case 14:
			addPoint(3, SideBlack, g.Context(), &ret)
			addPoint(5, SideBlack, g.Context(), &ret)
			addPoint(23, SideBlack, g.Context(), &ret)
			addPoint(25, SideBlack, g.Context(), &ret)
		case 23:
			addPoint(14, SideBlack, g.Context(), &ret)
		case 25:
			addPoint(14, SideBlack, g.Context(), &ret)
		}
	}

	if g.Side() == SideBlack {
		switch point {
		case 93:
			addPoint(84, SideRed, g.Context(), &ret)
		case 95:
			addPoint(84, SideRed, g.Context(), &ret)
		case 84:
			addPoint(93, SideRed, g.Context(), &ret)
			addPoint(95, SideRed, g.Context(), &ret)
			addPoint(73, SideRed, g.Context(), &ret)
			addPoint(75, SideRed, g.Context(), &ret)
		case 73:
			addPoint(84, SideRed, g.Context(), &ret)
		case 75:
			addPoint(84, SideRed, g.Context(), &ret)
		}
	}

	return ret
}
