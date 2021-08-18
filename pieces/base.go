package pieces

const (
	Xmin = 0
	Xmax = 8
	Ymin = 0
	Ymax = 9
)

func addPoint(point int, side Side, points *[]*Point) {
	if piece, ok := Metrics[point]; ok {
		if piece.Side() == side {
			*points = append(*points, NewPoint(point, piece))
		}
	} else {
		*points = append(*points, NewPoint(point, nil))
	}
}

type Side int

const (
	SideRed   Side = 1
	SideBlack Side = -1
)

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

var (
	Red     map[int]Piece
	Black   map[int]Piece
	Metrics map[int]Piece

	_ Piece = &Rook{}
	_ Piece = &Horse{}
	_ Piece = &Minister{}
	_ Piece = &Guard{}
	_ Piece = &King{}
	_ Piece = &Cannon{}
	_ Piece = &Pawn{}
)

type Piece interface {
	Side() Side
	Point() Point

	Next() []*Point
}

type PieceImpl struct {
	side  Side
	point Point
}

func (p *PieceImpl) Side() Side {
	return p.side
}

func (p *PieceImpl) Point() Point {
	return p.point
}
