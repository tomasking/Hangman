# Hangman

## Thoughts / explanations / todos

* I'm passing userId in the query string, this could easily be changed to interfere with other peoples games so wouldn't be secure. A better method could be to using JWT with the user id inside the token.
* I'm not stopping the user selecting a completed game