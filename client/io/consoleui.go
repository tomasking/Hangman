package io

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

type ConsoleUI struct{}

func NewConsoleUI() UI {
	uilayer := &ConsoleUI{}
	return uilayer
}

func (m ConsoleUI) DisplayGames(userGames []contracts.UserGame) {

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

func (m *ConsoleUI) SelectGame() int {
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

func (m *ConsoleUI) DisplayGameState(gameState model.GameState) {

	status := ""

	switch gameState.Status {
	case model.Completed:
		status = "(COMPLETED!)"
	case model.LastGoWasHit:
		status = "(HIT)"
	case model.LastGoWasMiss:
		status = "(MISS)"
	case model.AlreadyGuessedLetter:
		status = "(Already guessed this letter)"
	}

	for _, c := range gameState.GuessedWord {
		fmt.Print(string(c))
		fmt.Print(" ")
	}

	fmt.Println(status)
}

func (m *ConsoleUI) ReadGuess() rune {
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
