package restclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../../contracts"
)

func LoadUserGames(userId string) []contracts.UserGame {

	var userGames []contracts.UserGame

	body, err := getRequest("http://localhost:8000/usergame?user=" + userId)

	if err != nil {
		log.Fatal("Could not load request: ", err)
	} else {

		err = json.Unmarshal(body, &userGames)
		if err != nil {
			log.Fatal("Could not json decode user games", err)
		}
	}

	return userGames
}

func LoadNewGame(userId string) contracts.UserGame {

	var userGame contracts.UserGame

	body, err := getRequest("http://localhost:8000/usergame/new?user=" + userId)

	if err != nil {
		log.Fatal("Could not load request: ", err)
	} else {

		err = json.Unmarshal(body, &userGame)
		if err != nil {
			log.Fatal("Could not json decode user game", err)
		}
	}
	return userGame
}

func LoadGame(userId string, gameId int) contracts.UserGame {

	var userGame contracts.UserGame

	url := "http://localhost:8000/usergame/" + strconv.Itoa(gameId) + "?user=" + userId
	body, err := getRequest(url)

	if err != nil {
		log.Fatal("Could not load request: ", err)
	} else {

		err = json.Unmarshal(body, &userGame)
		if err != nil {
			log.Fatal("Could not json decode user game", err)
		}
	}

	return userGame
}

func getRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return nil, err2
	}
	return body, nil
}

func UpdateUserGame(userId string, gameId int, guesses []string, gameCompleted bool, word string) {

	var userGame = contracts.UpdateUserGame{Guesses: guesses, Completed: gameCompleted, Word: word}

	contents, _ := json.Marshal(userGame)

	body := bytes.NewBufferString(string(contents))

	url := "http://localhost:8000/usergame/" + strconv.Itoa(gameId) + "?user=" + userId

	rsp, err := http.Post(url, "application/json", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	_, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
}
