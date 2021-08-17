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
	piece Piece
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
	return &Point{}
}

var (
	Red     map[int]Piece
	Black   map[int]Piece
	Metrics map[int]Piece
)

type Piece interface {
	Side() Side
	Coordination() Point

	Next() []Point
	Aim() []Piece
}

type Rook struct {
}

type Minister struct {
}

type Guard struct {
}

type Cannon struct {
}

type Pawn struct {
}
