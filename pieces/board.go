package pieces

type Board struct {
	Red     map[int]Piece
	Black   map[int]Piece
	Metrics map[int]Piece
}
