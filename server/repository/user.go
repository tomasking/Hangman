package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"../../contracts"
)

const baseUserDir = "./repository/userdb/"

func LoadUserGames(userId string) []contracts.UserGame {

	var userGames []contracts.UserGame
	filename := baseUserDir + userId + ".txt"

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		contents, _ := json.Marshal(userGames)
		ioutil.WriteFile(filename, []byte(contents), 0644)
	} else {

		err = json.Unmarshal(dat, &userGames)
		if err != nil {
			log.Fatal("Could not load user file")
		}
	}

	return userGames
}

func LoadUserGame(userId string, gameId int) contracts.UserGame {

	userGames := LoadUserGames(userId)

	for _, userGame := range userGames {

		if gameId == userGame.GameId {
			return userGame
			break
		}

	}

	log.Fatal("Could not find game Id ", gameId)
	return contracts.UserGame{} //TODO: null?
}

func UpdateUserGame(userId string, gameId int, gameState contracts.UpdateUserGame) error {

	userGames := LoadUserGames(userId)

	found := false

	for index, userGame := range userGames {
		if gameId == userGame.GameId {
			userGame.GameId = gameId
			userGame.Guesses = gameState.Guesses
			userGame.Completed = gameState.Completed
			userGame.Word = gameState.Word
			userGames[index] = userGame
			found = true
			break
		}
	}

	if !found {
		userGame := contracts.UserGame{GameId: gameId, Guesses: gameState.Guesses, Completed: false, Word: gameState.Word}
		userGames = append(userGames, userGame)
	}

	filename := baseUserDir + userId + ".txt"
	contents, _ := json.Marshal(userGames)
	ioutil.WriteFile(filename, []byte(contents), 0644)

	return nil //TODO error handling
}
