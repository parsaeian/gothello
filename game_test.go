package othello

import (
	"testing"
)

const initBoard = `
  A B C D E F G H
1 - - - - - - - -
2 - - - - - - - -
3 - - - - - - - -
4 - - - # * - - -
5 - - - * # - - -
6 - - - - - - - -
7 - - - - - - - -
8 - - - - - - - -
`

func TestGame(t *testing.T) {
	t.Run("create new game", func(t *testing.T) {
		_, err := NewGame()

		if err != nil {
			t.Errorf("can not create new game!, %v", err)
		}
	})

	t.Run("show new game's board", func(t *testing.T) {
		game, _ := NewGame()

		got, err := game.board.Draw()

		if err != nil {
			t.Fatalf("error in draw new game's board, %v!", err)
		}

		want := initBoard

		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
