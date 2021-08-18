package pieces

type Pawn struct {
	PieceImpl
}

func (p *Pawn) Next() []*Point {
	ret := make([]*Point, 0)

	x0 := p.Point().X()
	y0 := p.Point().Y()
	point := p.Point().Point()

	if p.Side() == SideRed {
		if y0 > 4 {
			if x0 > Xmin {
				addPoint(point-1, SideBlack, &ret)
			}

			if x0 < Xmax {
				addPoint(point+1, SideBlack, &ret)
			}

			if y0 < Ymax {
				addPoint(point+10, SideBlack, &ret)
			}
		} else {
			addPoint(point+10, SideBlack, &ret)
		}
	}

	if p.Side() == SideBlack {
		if y0 < 5 {
			if x0 > Xmin {
				addPoint(point-1, SideRed, &ret)
			}

			if x0 < Xmax {
				addPoint(point+1, SideRed, &ret)
			}

			if y0 > Ymin {
				addPoint(point-10, SideRed, &ret)
			}
		} else {
			addPoint(point-10, SideRed, &ret)
		}
	}

	return ret
}
