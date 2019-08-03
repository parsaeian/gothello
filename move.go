package othello

import "fmt"

type Square uint8 // A Square is othello board square (between 0-63)

// MoveStr decodes the given string in game's notation
// and calls the Move function.
func (g *Game) MoveStr(s string) error {
	square, err := g.notation.Decode(s)
	if err != nil {
		return fmt.Errorf("error in decode string, %v", err)
	}
	bs := g.board.blackSquares
	g.board.blackSquares = bs | 1<<square | 1<<36
	g.board.whiteSquares = 1 << 27
	g.turn = whiteTurn
	return nil
}
