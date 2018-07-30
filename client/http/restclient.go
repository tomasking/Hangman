package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../../contracts"
)

type RestClient struct{}

func NewRestClient() RestLayer {
	return &RestClient{}
}

const baseUrl = "http://localhost:8000/usergame" // this should come from config

func (m RestClient) LoadUserGames(userId string) []contracts.UserGame {

	var userGames []contracts.UserGame

	body, err := getRequest(baseUrl + "?user=" + userId)

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

func (m RestClient) LoadNewGame(userId string) contracts.UserGame {

	var userGame contracts.UserGame

	body, err := getRequest(baseUrl + "/new?user=" + userId)

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

func (m RestClient) LoadGame(userId string, gameId int) contracts.UserGame {

	var userGame contracts.UserGame

	url := baseUrl + "/" + strconv.Itoa(gameId) + "?user=" + userId
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

func (m RestClient) UpdateUserGame(userId string, gameId int, guesses []string, gameCompleted bool, word string) {

	var userGame = contracts.UpdateUserGame{Guesses: guesses, Completed: gameCompleted, Word: word}

	contents, _ := json.Marshal(userGame)

	body := bytes.NewBufferString(string(contents))

	url := baseUrl + "/" + strconv.Itoa(gameId) + "?user=" + userId

	rsp, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Println(err)
		return // ignore error allowing user to carry on, hopefully it works next time
	}
	defer rsp.Body.Close()
}
