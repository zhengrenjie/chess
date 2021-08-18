package pieces

type Side int
type PieceType int

const (
	Xmin = 0
	Xmax = 8
	Ymin = 0
	Ymax = 9

	SideRed   Side = 1
	SideBlack Side = -1

	TypeRook PieceType = iota + 1
	TypeHorse
	TypeMinister
	TypeGuard
	TypeKing
	TypeCannon
	TypePawn
)

var (
	_ Piece = &Rook{}
	_ Piece = &Horse{}
	_ Piece = &Minister{}
	_ Piece = &Guard{}
	_ Piece = &King{}
	_ Piece = &Cannon{}
	_ Piece = &Pawn{}
)

func addPoint(point int, side Side, board *Board, points *[]*Point) {
	if piece, ok := board.Metrics[point]; ok {
		if piece.Side() == side {
			*points = append(*points, NewPoint(point, piece))
		}
	} else {
		*points = append(*points, NewPoint(point, nil))
	}
}

type Point struct {
	point int
	x     int
	y     int
	piece Piece // target, could be killed.
}

func (p Point) Piece() Piece {
	return p.piece
}

func (p Point) Point() int {
	return p.point
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}

func NewPoint(point int, piece Piece) *Point {
	return &Point{
		point: point,
		x:     point % 10,
		y:     point / 10,
		piece: piece,
	}
}

type Piece interface {
	Context() *Board
	Side() Side
	Point() Point

	Next() []*Point
}

type PieceImpl struct {
	side  Side
	point Point
	board *Board
}

func (p *PieceImpl) Side() Side {
	return p.side
}

func (p *PieceImpl) Point() Point {
	return p.point
}

func (p *PieceImpl) Context() *Board {
	return p.board
}
