# Hangman

## Thoughts / explanations / todos

* I'm passing userId in the query string, this could easily be changed to interfere with other peoples games so wouldn't be secure. A better method could be to using JWT with the user id inside the token.
* I've also harcoded the userId for now to save time
* I'm not stopping the user selecting a completed game
* I haven't really validated the user input either, I should only be accepting a lower case letter between a-z 


## Error Handling

* I've hardly done any error handling, if the client has trouble communicating with the server than I'm often just logging fatal. Ideally, I would like ther server to return an error object containing a friendly error message to display to the user. In the client I could then have retry logic to retry the call.
* After every guess I update the server with the current guesses for a game. In this case it doesn't matter if communication goes down for a bit then comes up later as I'm sending all the guesses every go.


## Testing

* Unfortuantely I didn't have time to do the tests (which in C# I would be doing TDD). I've got around to working out interfaces so in the client I put interfaces in front of the Console input/output and the http calls to the server. This meant I would have been able to do testing around the business logic mocking out these adapters. In the server it basically is a lot of crud calls so I would normally have integration tests around these covering the http calls and saving to a backing store (in this case I cheated and saved to file).