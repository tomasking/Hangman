# Hangman

## Todos

* I'm passing userId in the query string, this could easily be changed to interfere with other peoples games so wouldn't be secure. A better method could be to using JWT with the user id inside the token.
* I've also harcoded the userId for now to save time
* I'm not stopping the user selecting a completed game
* I haven't really validated the user input either, I should only be accepting a lower case letter between a-z 
* I've hardly done any error handling, if the client has trouble communicating with the server than I'm often just logging fatal. Ideally, I would like ther server to return an error object containing a friendly error message to display to the user. In the client I could then have retry logic to retry the call.
* Need to add tests
