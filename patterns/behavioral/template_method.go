package behavioral

import "fmt"

// Template Method is a behavioral design pattern that defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.

// https://refactoring.guru/design-patterns/template-method

// 1. Common Template Method
type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
	Timeline() []string
}

type Tournament struct {
	Game
}

func NewTournament(g Game) *Tournament {
	return &Tournament{g}
}

func (t *Tournament) PlayGame() {
	t.Game.Start()
	for !t.Game.HaveWinner() {
		t.Game.TakeTurn()
	}
	fmt.Printf("Player %d wins.\n", t.Game.WinningPlayer())
}

func (t *Tournament) Timeline() []string {
	return t.Game.Timeline()
}

type chess struct {
	turn, maxTurns, currentPlayer int
	timeline                      []string
}

func NewGameOfChess() Game {
	return &chess{
		turn:          1,
		maxTurns:      10,
		currentPlayer: 0,
	}
}

func (c *chess) Start() {
	c.timeline = append(c.timeline, "Starting a new game of chess.")
}

func (c *chess) TakeTurn() {
	c.turn++
	c.timeline = append(c.timeline, fmt.Sprintf("Turn %d taken by player %d\n", c.turn, c.currentPlayer))
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	c.timeline = append(c.timeline, fmt.Sprintf("Player %d won the game", c.currentPlayer))
	return c.currentPlayer
}

func (c *chess) Timeline() []string {
	return c.timeline
}

// 2. Function Template Method
func PlayGame(
	start, takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int,
) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}
