package main

import (
	"./hangman"
	"./http"
	"./io"
	"./localstorage"
)

func main() {

	userId := localstorage.LoadExistingUser()

	consoleUi := io.NewConsoleUI()
	restClient := http.NewRestClient()

	for {
		game := hangman.SelectGame(consoleUi, restClient, userId)
		hangman.PlayGame(consoleUi, restClient, userId, game)
	}
}
