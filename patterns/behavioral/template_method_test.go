package behavioral_test

import (
	"fmt"
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestTemplateMethod(t *testing.T) {
	t.Run("Should implement common template method", func(t *testing.T) {
		g := behavioral.NewGameOfChess()
		tour := behavioral.NewTournament(g)

		tour.PlayGame()

		assert.Len(t, tour.Timeline(), 11)
	})

	t.Run("Should implement functional template method", func(t *testing.T) {
		turn, maxTurns, currentPlayer := 1, 10, 0
		timeline := []string{}

		start := func() {
			timeline = append(timeline, "Starting a new game of chess.")
		}
		takeTurn := func() {
			turn++
			timeline = append(timeline, fmt.Sprintf("Turn %d taken by player %d\n", turn, currentPlayer))
			currentPlayer = 1 - currentPlayer
		}

		haveWinner := func() bool {
			return turn == maxTurns
		}

		winningPlayer := func() int {
			timeline = append(timeline, fmt.Sprintf("Player %d won the game", currentPlayer))
			return currentPlayer
		}

		behavioral.PlayGame(start, takeTurn, haveWinner, winningPlayer)

		assert.Len(t, timeline, 11)
	})
}
