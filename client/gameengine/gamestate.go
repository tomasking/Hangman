package gameengine

type GameStatus int

const (
	NewGame              GameStatus = 0
	AlreadyGuessedLetter GameStatus = 1
	LastGoWasHit         GameStatus = 2
	LastGoWasMiss        GameStatus = 3
	Completed            GameStatus = 4
)

type GameState struct {
	GameId      int
	Word        string
	Guesses     []string
	GuessedWord []rune
	Status      GameStatus
}
