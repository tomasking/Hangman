package contracts

type UpdateUserGame struct {
	Guesses   []string `json:"guesses"`
	Completed bool     `json:"completed"`
	Word      string   `json:"word"` // you wouldn't need to send this, I'm just doing it now to keep dataaccess simpler
}
