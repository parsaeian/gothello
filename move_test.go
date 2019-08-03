package othello

import "testing"

const firstMoveNewGameBoard = `
  A B C D E F G H
1 - - - - - - - -
2 - - - - - - - -
3 - - - - - - - -
4 - - - # * - - -
5 - - - * * * - -
6 - - - - - - - -
7 - - - - - - - -
8 - - - - - - - -
`

func TestMove(t *testing.T) {
	t.Run("do first move in new game", func(t *testing.T) {
		game, _ := NewGame()
		err := game.MoveStr("f5")

		if err != nil {
			t.Errorf("error in moveStr, %v", err)
		}

		if game.turn != whiteTurn {
			t.Error("error in update turn")
		}

		got, err := game.board.Draw()
		if err != nil {
			t.Fatalf("error in draw new game's board, %v!", err)
		}
		want := firstMoveNewGameBoard
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
