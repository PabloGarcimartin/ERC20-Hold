# ERC20-Hold
ERC20 token supporting balances on hold

ERC20-HOLDABLE tests:
1-On a cmd go to the project folder and run: "truffle develop"
2-Run: "test" -> 11 tests on the ERC20-Holdable token should pass

BETTING-GAME test:
1-On a cmd go to the project folder and run: "truffle develop"
2-Run: "migrate" and wait for the migration to finish.
3-On a new cmd got o the project folder and run: "node index.js" and wait for "API escuchando en el puerto 8080" message
4-Import the BETGAME.postman_collection in postman and run 4 times the 'Add player' request.
5-Run all the 'bet' requests, for each one a message of 'bet created' should appear on the api cmd.
6-On the 4th bet, we should see the winner on the api cmd.
7-Run the 'Get user balance and bets' requests to check the users balances after the bet was executed.





