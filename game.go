package othello

import (
	"fmt"
	"strconv"
)

type Board struct {
	blackSquares uint64
	whiteSquares uint64
}

type Game struct {
	p1    *Player
	p2    *Player
	board *Board
}

func NewBoard() (b *Board) {
	b = &Board{
		blackSquares: 1<<28 | 1<<35,
		whiteSquares: 1<<27 | 1<<36}
	return
}

func NewGame(p1, p2 *Player) (Game, error) {
	if p1 == nil {
		return Game{}, fmt.Errorf("player1 can't be nil")
	}
	if p2 == nil {
		return Game{}, fmt.Errorf("player1 can't be nil")
	}

	board := NewBoard()
	return Game{p1, p2, board}, nil
}

// Draw returns visual representation of the Board useful for debugging.
func (b *Board) Draw() (string, error) {
	s := "\n  A B C D E F G H\n"
	for i := 0; i <= 7; i++ {
		s += strconv.FormatInt(int64(1+i), 10)
		for j := 0; j <= 7; j++ {
			squareNumber := uint64(i*8 + j)
			isBlack := 1 << squareNumber & b.blackSquares
			isWhite := 1 << squareNumber & b.whiteSquares

			if isBlack != 0 && isWhite != 0 {
				return "", fmt.Errorf("error in draw board, cell %d", squareNumber)
			}
			if isBlack != 0 {
				s += " *"
			} else if isWhite != 0 {
				s += " #"
			} else {
				s += " -"
			}

		}
		s += "\n"
	}

	return s, nil
}
