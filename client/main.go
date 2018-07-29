package main

import (
	"../contracts"
	"./gameengine"
	"./localstorage"
	"./restclient"
	"./ui"
)

func main() {

	userId := localstorage.LoadExistingUser()

	for {
		game := selectGame(userId)
		playGame(userId, game)
	}
}

func selectGame(userId string) contracts.UserGame {
	games := restclient.LoadUserGames(userId)
	ui.DisplayGames(games)
	gameId := ui.SelectGame()

	var game contracts.UserGame

	if gameId == 0 {
		game = restclient.LoadNewGame(userId)
	} else {
		game = restclient.LoadGame(userId, gameId)
	}
	return game
}

func playGame(userId string, game contracts.UserGame) {
	gameState := gameengine.Initialize(game)
	ui.DisplayGameState(gameState)

	for {
		guess := ui.ReadGuess()
		gameState = gameengine.GuessLetter(gameState, guess)
		ui.DisplayGameState(gameState)

		gameCompleted := gameState.Status == gameengine.Completed

		restclient.UpdateUserGame(userId, gameState.GameId, gameState.Guesses, gameCompleted, gameState.Word)

		if gameCompleted {
			break
		}
	}
}
