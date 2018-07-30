package http

import (
	"../../contracts"
)

type RestLayer interface {
	LoadUserGames(userId string) []contracts.UserGame
	LoadNewGame(userId string) contracts.UserGame
	LoadGame(userId string, gameId int) contracts.UserGame
	UpdateUserGame(userId string, gameId int, guesses []string, gameCompleted bool, word string)
}
