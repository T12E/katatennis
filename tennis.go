package main

import (
	"math"
)

var scores = []string{"00", "15", "30", "40"}

type Game struct {
	playerOneScore int
	playerTwoScore int
	playerOneName  string
	playerTwoName  string
}

func NewGame(playerOneName, playerTwoName string) *Game {
	return &Game{playerOneName: playerOneName, playerTwoName: playerTwoName}
}

func computeScore(score int, c chan string) {
	c <- scores[score]
}

func (g *Game) Score() string {
	if g.playerOneScore < 4 && g.playerTwoScore < 4 {
		c1 := make(chan string)
		c2 := make(chan string)
		go computeScore(g.playerOneScore, c1)
		go computeScore(g.playerTwoScore, c2)
		return <-c1 + "-" + <-c2
	}
	switch g.getPointsDiff() {
	case 0:
		return "Deuce"
	case 1:
		return "ADV " + g.leadingPlayer()
	}
	return g.leadingPlayer() + " win"
}

func (g *Game) getPointsDiff() int {
	return int(math.Abs(float64(g.playerOneScore - g.playerTwoScore)))
}

func (g *Game) leadingPlayer() string {
	if g.playerOneScore > g.playerTwoScore {
		return g.playerOneName
	} else {
		return g.playerTwoName
	}
}

func (g *Game) PlayerOneScores() {
	g.playerOneScore++
}

func (g *Game) PlayerTwoScores() {
	g.playerTwoScore++
}
