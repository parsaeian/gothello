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
		player1 := Player{}
		player2 := Player{}
		_, err := NewGame(&player1, &player2)

		if err != nil {
			t.Errorf("can not create new game!, %d", err)
		}
	})

	t.Run("show new game's board", func(t *testing.T) {
		player1 := Player{}
		player2 := Player{}
		game, _ := NewGame(&player1, &player2)

		got, err := game.board.Draw()

		if err != nil {
			t.Fatalf("erro in create new game's board, %d!", err)
		}

		want := initBoard

		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	})
}
