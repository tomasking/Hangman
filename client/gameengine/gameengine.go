package gameengine

import (
	"../../contracts"
)

func Initialize(game contracts.UserGame) GameState {

	var guessedWord []rune
	for _, letter := range game.Word {
		if contains(game.Guesses, letter) {
			guessedWord = append(guessedWord, letter)
		} else {
			guessedWord = append(guessedWord, '_')
		}
	}

	currentState := GameState{GameId: game.GameId, Word: game.Word, Guesses: game.Guesses, GuessedWord: guessedWord, Status: NewGame}

	return currentState
}

func GuessLetter(currentState GameState, letter rune) GameState {

	if contains(currentState.Guesses, letter) {
		currentState.Status = AlreadyGuessedLetter
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
		currentState.Status = Completed
	} else if hit {
		currentState.Status = LastGoWasHit
	} else {
		currentState.Status = LastGoWasMiss
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
