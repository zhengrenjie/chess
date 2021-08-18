package pieces

type Minister struct {
	PieceImpl
}

func (m *Minister) Next() []*Point {
	ret := make([]*Point, 0)
	point := m.Point().Point()

	if m.Side() == SideRed {
		switch point {
		case 2:
			addPoint(20, SideBlack, &ret)
			addPoint(24, SideBlack, &ret)
		case 6:
			addPoint(24, SideBlack, &ret)
			addPoint(28, SideBlack, &ret)
		case 20:
			addPoint(2, SideBlack, &ret)
			addPoint(42, SideBlack, &ret)
		case 24:
			addPoint(2, SideBlack, &ret)
			addPoint(42, SideBlack, &ret)
			addPoint(6, SideBlack, &ret)
			addPoint(46, SideBlack, &ret)
		case 28:
			addPoint(6, SideBlack, &ret)
			addPoint(46, SideBlack, &ret)
		case 42:
			addPoint(20, SideBlack, &ret)
			addPoint(24, SideBlack, &ret)
		case 46:
			addPoint(24, SideBlack, &ret)
			addPoint(28, SideBlack, &ret)
		}
	}

	if m.Side() == SideBlack {
		switch point {
		case 92:
			addPoint(70, SideBlack, &ret)
			addPoint(74, SideBlack, &ret)
		case 96:
			addPoint(74, SideBlack, &ret)
			addPoint(78, SideBlack, &ret)
		case 70:
			addPoint(92, SideBlack, &ret)
			addPoint(52, SideBlack, &ret)
		case 74:
			addPoint(92, SideBlack, &ret)
			addPoint(52, SideBlack, &ret)
			addPoint(96, SideBlack, &ret)
			addPoint(56, SideBlack, &ret)
		case 78:
			addPoint(96, SideBlack, &ret)
			addPoint(56, SideBlack, &ret)
		case 52:
			addPoint(70, SideBlack, &ret)
			addPoint(74, SideBlack, &ret)
		case 56:
			addPoint(74, SideBlack, &ret)
			addPoint(78, SideBlack, &ret)
		}
	}

	return ret
}
