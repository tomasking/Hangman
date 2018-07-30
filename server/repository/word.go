package repository

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

var words []string

func LoadWord() (string, int) {

	if words == nil {
		words = loadFile()
	}

	gamenumber := rand.Intn(851) //The random selection of new games doesn't seem to be that random
	word := strings.ToLower(words[gamenumber])
	return word, gamenumber
}

func loadFile() []string {

	dat, err := ioutil.ReadFile("./repository/worddb/words.txt")
	if err != nil {
		fmt.Println("Could not load file: ", err)
		return nil
	}

	return strings.Split(string(dat), "\n")
}
