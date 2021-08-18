package pieces

type Board struct {
	Red     map[int]Piece
	Black   map[int]Piece
	Metrics map[int]Piece
}

func NewBoard() *Board {
	ret := &Board{
		Red:     make(map[int]Piece),
		Black:   make(map[int]Piece),
		Metrics: make(map[int]Piece),
	}

	/* red */
	ret.Metrics[0] = &Rook{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(0, nil), board: ret}}
	ret.Metrics[8] = &Rook{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(8, nil), board: ret}}
	ret.Metrics[1] = &Horse{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(1, nil), board: ret}}
	ret.Metrics[7] = &Horse{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(7, nil), board: ret}}
	ret.Metrics[2] = &Minister{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(2, nil), board: ret}}
	ret.Metrics[6] = &Minister{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(6, nil), board: ret}}
	ret.Metrics[3] = &Guard{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(3, nil), board: ret}}
	ret.Metrics[5] = &Guard{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(5, nil), board: ret}}
	ret.Metrics[4] = &King{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(4, nil), board: ret}}
	ret.Metrics[21] = &Cannon{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(21, nil), board: ret}}
	ret.Metrics[27] = &Cannon{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(27, nil), board: ret}}
	ret.Metrics[30] = &Pawn{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(30, nil), board: ret}}
	ret.Metrics[32] = &Pawn{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(32, nil), board: ret}}
	ret.Metrics[34] = &Pawn{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(34, nil), board: ret}}
	ret.Metrics[36] = &Pawn{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(36, nil), board: ret}}
	ret.Metrics[38] = &Pawn{PieceImpl: PieceImpl{side: SideRed, point: NewPoint(38, nil), board: ret}}

	/* black */
	ret.Metrics[90] = &Rook{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(0, nil), board: ret}}
	ret.Metrics[98] = &Rook{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(8, nil), board: ret}}
	ret.Metrics[91] = &Horse{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(1, nil), board: ret}}
	ret.Metrics[97] = &Horse{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(7, nil), board: ret}}
	ret.Metrics[92] = &Minister{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(2, nil), board: ret}}
	ret.Metrics[96] = &Minister{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(6, nil), board: ret}}
	ret.Metrics[93] = &Guard{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(3, nil), board: ret}}
	ret.Metrics[95] = &Guard{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(5, nil), board: ret}}
	ret.Metrics[94] = &King{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(4, nil), board: ret}}
	ret.Metrics[71] = &Cannon{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(21, nil), board: ret}}
	ret.Metrics[77] = &Cannon{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(27, nil), board: ret}}
	ret.Metrics[60] = &Pawn{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(30, nil), board: ret}}
	ret.Metrics[62] = &Pawn{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(32, nil), board: ret}}
	ret.Metrics[64] = &Pawn{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(34, nil), board: ret}}
	ret.Metrics[66] = &Pawn{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(36, nil), board: ret}}
	ret.Metrics[68] = &Pawn{PieceImpl: PieceImpl{side: SideBlack, point: NewPoint(38, nil), board: ret}}

	for k, v := range ret.Metrics {
		switch v.Side() {
		case SideRed:
			ret.Red[k] = v
		case SideBlack:
			ret.Black[k] = v
		}
	}

	return ret
}
