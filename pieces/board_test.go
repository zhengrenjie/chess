package pieces

import (
	"fmt"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	for j := 0; j <= Ymax; j++ {
		for i := 0; i <= Xmax; i++ {
			point := j*10 + i
			if piece, ok := board.Metrics[point]; ok {
				fmt.Print(piece)
			} else {
				fmt.Print("\u3000")
			}
		}
		fmt.Print("\n")
	}
}
