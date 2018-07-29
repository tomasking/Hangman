package contracts

type UserGame struct {
	GameId    int      `json:"gameId"`
	Word      string   `json:"word"`
	Guesses   []string `json:"guesses"`
	Completed bool     `json:"completed"`
}
