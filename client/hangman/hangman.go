package hangman

import (
	"../../contracts"
	"../http"
	"../io"
	"../model"
)

func SelectGame(ui io.UI, restclient http.RestLayer, userId string) contracts.UserGame {
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

func PlayGame(ui io.UI, restclient http.RestLayer, userId string, game contracts.UserGame) {
	gameState := Initialize(game)
	ui.DisplayGameState(gameState)

	for {
		guess := ui.ReadGuess()
		gameState = GuessLetter(gameState, guess)
		ui.DisplayGameState(gameState)

		gameCompleted := gameState.Status == model.Completed

		restclient.UpdateUserGame(userId, gameState.GameId, gameState.Guesses, gameCompleted, gameState.Word)

		if gameCompleted {
			break
		}
	}
}
