package pieces

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
)

type Piece interface {
	Side() Side
	Point() Point

	Next() []Point
}

type Rook struct {
}

type Minister struct {
}

type Guard struct {
}

type Pawn struct {
}
