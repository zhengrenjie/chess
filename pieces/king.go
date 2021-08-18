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
			addPoint(4, SideBlack, &ret)
			addPoint(13, SideBlack, &ret)
		case 4:
			addPoint(3, SideBlack, &ret)
			addPoint(5, SideBlack, &ret)
			addPoint(14, SideBlack, &ret)
		case 5:
			addPoint(4, SideBlack, &ret)
			addPoint(15, SideBlack, &ret)
		case 13:
			addPoint(3, SideBlack, &ret)
			addPoint(23, SideBlack, &ret)
			addPoint(14, SideBlack, &ret)
		case 14:
			addPoint(13, SideBlack, &ret)
			addPoint(15, SideBlack, &ret)
			addPoint(4, SideBlack, &ret)
			addPoint(24, SideBlack, &ret)
		case 15:
			addPoint(5, SideBlack, &ret)
			addPoint(25, SideBlack, &ret)
			addPoint(14, SideBlack, &ret)
		case 23:
			addPoint(13, SideBlack, &ret)
			addPoint(24, SideBlack, &ret)
		case 24:
			addPoint(23, SideBlack, &ret)
			addPoint(25, SideBlack, &ret)
			addPoint(14, SideBlack, &ret)
		case 25:
			addPoint(15, SideBlack, &ret)
			addPoint(24, SideBlack, &ret)
		}
	}

	if k.Side() == SideBlack {
		switch k.Point().Point() {
		case 93:
			addPoint(94, SideRed, &ret)
			addPoint(83, SideRed, &ret)
		case 94:
			addPoint(93, SideRed, &ret)
			addPoint(95, SideRed, &ret)
			addPoint(84, SideRed, &ret)
		case 95:
			addPoint(94, SideRed, &ret)
			addPoint(85, SideRed, &ret)
		case 83:
			addPoint(93, SideRed, &ret)
			addPoint(73, SideRed, &ret)
			addPoint(84, SideRed, &ret)
		case 84:
			addPoint(83, SideRed, &ret)
			addPoint(85, SideRed, &ret)
			addPoint(94, SideRed, &ret)
			addPoint(74, SideRed, &ret)
		case 85:
			addPoint(95, SideRed, &ret)
			addPoint(75, SideRed, &ret)
			addPoint(84, SideRed, &ret)
		case 73:
			addPoint(83, SideRed, &ret)
			addPoint(74, SideRed, &ret)
		case 74:
			addPoint(73, SideRed, &ret)
			addPoint(75, SideRed, &ret)
			addPoint(84, SideRed, &ret)
		case 75:
			addPoint(85, SideRed, &ret)
			addPoint(74, SideRed, &ret)
		}
	}

	return ret
}
