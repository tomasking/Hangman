package hangman

import (
	"../../contracts"
	"../model"
)

func Initialize(game contracts.UserGame) model.GameState {

	var guessedWord []rune
	for _, letter := range game.Word {
		if contains(game.Guesses, letter) {
			guessedWord = append(guessedWord, letter)
		} else {
			guessedWord = append(guessedWord, '_')
		}
	}

	currentState := model.GameState{GameId: game.GameId, Word: game.Word, Guesses: game.Guesses, GuessedWord: guessedWord, Status: model.NewGame}

	return currentState
}

func GuessLetter(currentState model.GameState, letter rune) model.GameState {

	if contains(currentState.Guesses, letter) {
		currentState.Status = model.AlreadyGuessedLetter
		return currentState
	}

	hit := false
	for i, c := range currentState.Word {
		if letter == c {
			currentState.GuessedWord[i] = c
			hit = true
		}
	}

	if string(currentState.GuessedWord) == currentState.Word {
		currentState.Status = model.Completed
	} else if hit {
		currentState.Status = model.LastGoWasHit
	} else {
		currentState.Status = model.LastGoWasMiss
	}

	currentState.Guesses = append(currentState.Guesses, string(letter))

	return currentState
}

func contains(letters []string, letter rune) bool {

	for _, currentLetter := range letters {
		if currentLetter == string(letter) {
			return true
		}
	}
	return false
}
