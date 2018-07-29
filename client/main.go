package main

import (
	"./hangman"
	"./localstorage"
)

func main() {

	userId := localstorage.LoadExistingUser()

	for {
		game := hangman.SelectGame(userId)
		hangman.PlayGame(userId, game)
	}
}
