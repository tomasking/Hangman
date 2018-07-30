package io

import (
	"../../contracts"
	"../model"
)

type UI interface {
	DisplayGames(userGames []contracts.UserGame)
	SelectGame() int
	DisplayGameState(gameState model.GameState)
	ReadGuess() rune
}
