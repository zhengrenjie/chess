package pieces

type Pawn struct {
	PieceImpl
}

func (p *Pawn) Type() PieceType {
	return TypePawn
}

func (p *Pawn) String() string {
	if p.Side() == SideRed {
		return "兵"
	}
	return "卒"
}

func (p *Pawn) Next() []*Point {
	ret := make([]*Point, 0)

	x0 := p.Point().X()
	y0 := p.Point().Y()
	point := p.Point().Point()

	if p.Side() == SideRed {
		if y0 > 4 {
			if x0 > Xmin {
				addPoint(point-1, SideBlack, p.Context(), &ret)
			}

			if x0 < Xmax {
				addPoint(point+1, SideBlack, p.Context(), &ret)
			}

			if y0 < Ymax {
				addPoint(point+10, SideBlack, p.Context(), &ret)
			}
		} else {
			addPoint(point+10, SideBlack, p.Context(), &ret)
		}
	}

	if p.Side() == SideBlack {
		if y0 < 5 {
			if x0 > Xmin {
				addPoint(point-1, SideRed, p.Context(), &ret)
			}

			if x0 < Xmax {
				addPoint(point+1, SideRed, p.Context(), &ret)
			}

			if y0 > Ymin {
				addPoint(point-10, SideRed, p.Context(), &ret)
			}
		} else {
			addPoint(point-10, SideRed, p.Context(), &ret)
		}
	}

	return ret
}
