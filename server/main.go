package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../contracts"
	"./repository"
	"github.com/gorilla/mux"
)

const port = "8000"

func main() {

	fmt.Println("Starting server on port:", port)
	setupWebserver()
}

func setupWebserver() {
	router := mux.NewRouter()
	router.PathPrefix("/Server")

	router.HandleFunc("/usergame", GetUserGames).Queries("user", "{userId}").Methods("GET")
	router.HandleFunc("/usergame/new", GetNewUserGame).Queries("user", "{userId}").Methods("GET")
	router.HandleFunc("/usergame/{id}", GetUserGame).Queries("user", "{userId}").Methods("GET")
	router.HandleFunc("/usergame/{id}", UpdateUserGame).Queries("user", "{userId}").Methods("POST") //it wasn't obvious how to do a put on the client so going with post for now

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetUserGames(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId := params["userId"]
	userGames := repository.LoadUserGames(userId)

	json.NewEncoder(w).Encode(userGames)
}

func GetNewUserGame(w http.ResponseWriter, r *http.Request) {

	word, id := repository.LoadWord()
	var game = contracts.UserGame{GameId: id, Word: word}
	json.NewEncoder(w).Encode(game)
}

func GetUserGame(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId := params["userId"]
	gameId, _ := strconv.Atoi(params["id"])
	userGame := repository.LoadUserGame(userId, gameId)

	json.NewEncoder(w).Encode(userGame)
}

func UpdateUserGame(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId := params["userId"]
	gameId, _ := strconv.Atoi(params["id"])

	var updateUserGame contracts.UpdateUserGame
	err := json.NewDecoder(r.Body).Decode(&updateUserGame)
	if err != nil {
		log.Fatal(err)
	}

	err = repository.UpdateUserGame(userId, gameId, updateUserGame)
	json.NewEncoder(w).Encode(err)
}
