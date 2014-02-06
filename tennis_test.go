package main

import "testing"

var (
	g             *Game
	playerOneName = "Ivan Lendl"
	playerTwoName = "John McEnroe"
)

func setup() {
	g = NewGame(playerOneName, playerTwoName)
}

func assertScore(t *testing.T, expected string) {
	s := g.Score()
	if s != expected {
		t.Errorf("Unexpected score %s, expected %s", s, expected)
	}
}

func createScore(player1, player2 int) {
	for i := 0; i < player1; i++ {
		g.PlayerOneScores()
	}
	for i := 0; i < player2; i++ {
		g.PlayerTwoScores()
	}
}

func TestNewGameReturns00(t *testing.T) {
	setup()

	assertScore(t, "00-00")
}

func TestPlayerOneScores(t *testing.T) {
	setup()

	createScore(1, 0)

	assertScore(t, "15-00")
}

func TestPlayerTwoScores(t *testing.T) {
	setup()

	createScore(0, 1)

	assertScore(t, "00-15")
}

func TestPlayerOneScores2(t *testing.T) {
	setup()

	createScore(2, 0)

	assertScore(t, "30-00")
}

func TestAll(t *testing.T) {
	setup()

	createScore(1, 1)

	assertScore(t, "15-15")
}

func TestAdvantagePlayerOne(t *testing.T) {
	setup()

	createScore(4, 3)

	assertScore(t, "ADV " + playerOneName)
}

func TestAdvantagePlayerTwo(t *testing.T) {
	setup()

	createScore(6, 7)

	assertScore(t, "ADV " + playerTwoName)
}

func TestDeuce(t *testing.T) {
	setup()

	createScore(4, 4)

	assertScore(t, "Deuce")
}

func TestDeuceAfterAdvantage(t *testing.T) {
	setup()

	createScore(8, 8)

	assertScore(t, "Deuce")
}

func TestVictoryPlayerOne(t *testing.T) {
	setup()

	createScore(4, 2)

	assertScore(t,playerOneName + " win")
}

func TestVictoryPlayerTwo(t *testing.T) {
	setup()

	createScore(5, 7)

	assertScore(t,playerTwoName + " win")
}
