package pieces

type Minister struct {
	PieceImpl
}

func (m *Minister) Type() PieceType {
	return TypeMinister
}

func (m *Minister) String() string {
	if m.Side() == SideRed {
		return "相"
	}
	return "象"
}

func (m *Minister) Next() []*Point {
	ret := make([]*Point, 0)
	point := m.Point().Point()

	if m.Side() == SideRed {
		switch point {
		case 2:
			addPoint(20, SideBlack, m.Context(), &ret)
			addPoint(24, SideBlack, m.Context(), &ret)
		case 6:
			addPoint(24, SideBlack, m.Context(), &ret)
			addPoint(28, SideBlack, m.Context(), &ret)
		case 20:
			addPoint(2, SideBlack, m.Context(), &ret)
			addPoint(42, SideBlack, m.Context(), &ret)
		case 24:
			addPoint(2, SideBlack, m.Context(), &ret)
			addPoint(42, SideBlack, m.Context(), &ret)
			addPoint(6, SideBlack, m.Context(), &ret)
			addPoint(46, SideBlack, m.Context(), &ret)
		case 28:
			addPoint(6, SideBlack, m.Context(), &ret)
			addPoint(46, SideBlack, m.Context(), &ret)
		case 42:
			addPoint(20, SideBlack, m.Context(), &ret)
			addPoint(24, SideBlack, m.Context(), &ret)
		case 46:
			addPoint(24, SideBlack, m.Context(), &ret)
			addPoint(28, SideBlack, m.Context(), &ret)
		}
	}

	if m.Side() == SideBlack {
		switch point {
		case 92:
			addPoint(70, SideBlack, m.Context(), &ret)
			addPoint(74, SideBlack, m.Context(), &ret)
		case 96:
			addPoint(74, SideBlack, m.Context(), &ret)
			addPoint(78, SideBlack, m.Context(), &ret)
		case 70:
			addPoint(92, SideBlack, m.Context(), &ret)
			addPoint(52, SideBlack, m.Context(), &ret)
		case 74:
			addPoint(92, SideBlack, m.Context(), &ret)
			addPoint(52, SideBlack, m.Context(), &ret)
			addPoint(96, SideBlack, m.Context(), &ret)
			addPoint(56, SideBlack, m.Context(), &ret)
		case 78:
			addPoint(96, SideBlack, m.Context(), &ret)
			addPoint(56, SideBlack, m.Context(), &ret)
		case 52:
			addPoint(70, SideBlack, m.Context(), &ret)
			addPoint(74, SideBlack, m.Context(), &ret)
		case 56:
			addPoint(74, SideBlack, m.Context(), &ret)
			addPoint(78, SideBlack, m.Context(), &ret)
		}
	}

	return ret
}
