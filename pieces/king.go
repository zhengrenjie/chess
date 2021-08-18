package pieces

const ()

type King struct {
	PieceImpl
}

func (k *King) Next() []*Point {
	ret := make([]*Point, 0)

	if k.Side() == SideRed {
		switch k.Point().Point() {
		case 3:
			addPoint(4, SideBlack, k.Context(), &ret)
			addPoint(13, SideBlack, k.Context(), &ret)
		case 4:
			addPoint(3, SideBlack, k.Context(), &ret)
			addPoint(5, SideBlack, k.Context(), &ret)
			addPoint(14, SideBlack, k.Context(), &ret)
		case 5:
			addPoint(4, SideBlack, k.Context(), &ret)
			addPoint(15, SideBlack, k.Context(), &ret)
		case 13:
			addPoint(3, SideBlack, k.Context(), &ret)
			addPoint(23, SideBlack, k.Context(), &ret)
			addPoint(14, SideBlack, k.Context(), &ret)
		case 14:
			addPoint(13, SideBlack, k.Context(), &ret)
			addPoint(15, SideBlack, k.Context(), &ret)
			addPoint(4, SideBlack, k.Context(), &ret)
			addPoint(24, SideBlack, k.Context(), &ret)
		case 15:
			addPoint(5, SideBlack, k.Context(), &ret)
			addPoint(25, SideBlack, k.Context(), &ret)
			addPoint(14, SideBlack, k.Context(), &ret)
		case 23:
			addPoint(13, SideBlack, k.Context(), &ret)
			addPoint(24, SideBlack, k.Context(), &ret)
		case 24:
			addPoint(23, SideBlack, k.Context(), &ret)
			addPoint(25, SideBlack, k.Context(), &ret)
			addPoint(14, SideBlack, k.Context(), &ret)
		case 25:
			addPoint(15, SideBlack, k.Context(), &ret)
			addPoint(24, SideBlack, k.Context(), &ret)
		}
	}

	if k.Side() == SideBlack {
		switch k.Point().Point() {
		case 93:
			addPoint(94, SideRed, k.Context(), &ret)
			addPoint(83, SideRed, k.Context(), &ret)
		case 94:
			addPoint(93, SideRed, k.Context(), &ret)
			addPoint(95, SideRed, k.Context(), &ret)
			addPoint(84, SideRed, k.Context(), &ret)
		case 95:
			addPoint(94, SideRed, k.Context(), &ret)
			addPoint(85, SideRed, k.Context(), &ret)
		case 83:
			addPoint(93, SideRed, k.Context(), &ret)
			addPoint(73, SideRed, k.Context(), &ret)
			addPoint(84, SideRed, k.Context(), &ret)
		case 84:
			addPoint(83, SideRed, k.Context(), &ret)
			addPoint(85, SideRed, k.Context(), &ret)
			addPoint(94, SideRed, k.Context(), &ret)
			addPoint(74, SideRed, k.Context(), &ret)
		case 85:
			addPoint(95, SideRed, k.Context(), &ret)
			addPoint(75, SideRed, k.Context(), &ret)
			addPoint(84, SideRed, k.Context(), &ret)
		case 73:
			addPoint(83, SideRed, k.Context(), &ret)
			addPoint(74, SideRed, k.Context(), &ret)
		case 74:
			addPoint(73, SideRed, k.Context(), &ret)
			addPoint(75, SideRed, k.Context(), &ret)
			addPoint(84, SideRed, k.Context(), &ret)
		case 75:
			addPoint(85, SideRed, k.Context(), &ret)
			addPoint(74, SideRed, k.Context(), &ret)
		}
	}

	return ret
}
