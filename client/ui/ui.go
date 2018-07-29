package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"../../contracts"
	"../model"
)

func DisplayGames(userGames []contracts.UserGame) {

	fmt.Println("")
	fmt.Println("Select Game")
	fmt.Println("===========")
	fmt.Println("0) NEW GAME")
	for _, game := range userGames {
		completed := ""

		var displayWord string
		if game.Completed {
			displayWord = game.Word
			completed = "(COMPLETED)"
		} else {

			var guessedWord []rune
			for _, letter := range game.Word {
				if contains(game.Guesses, letter) {
					guessedWord = append(guessedWord, letter)
				} else {
					guessedWord = append(guessedWord, '_')
				}
				guessedWord = append(guessedWord, ' ')
			}
			displayWord = string(guessedWord)
		}

		fmt.Printf("%d) %s %s", game.GameId, displayWord, completed)
		fmt.Println("")
	}
}

func SelectGame() int {
	fmt.Print("Select game: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", 1)
	// TODO: validation /error handling
	gameNumber, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	return gameNumber
}

func DisplayGameState(gameState model.GameState) {

	status := ""
	//TODO: Switch
	if gameState.Status == model.Completed {
		status = "(COMPLETED!)"
	} else if gameState.Status == model.LastGoWasHit {
		status = "(HIT)"
	} else if gameState.Status == model.LastGoWasMiss {
		status = "(MISS)"
	} else if gameState.Status == model.AlreadyGuessedLetter {
		status = "(Already guessed this letter)"
	}

	for _, c := range gameState.GuessedWord {
		fmt.Print(string(c))
		fmt.Print(" ")
	}

	fmt.Println(status)
}

func ReadGuess() rune {
	reader := bufio.NewReader(os.Stdin)

	guess, _, _ := reader.ReadRune()

	// TODO: validation on single letter character and lower case
	return guess
}

func contains(letters []string, letter rune) bool {

	for _, currentLetter := range letters {
		if currentLetter == string(letter) {
			return true
		}
	}
	return false
}
