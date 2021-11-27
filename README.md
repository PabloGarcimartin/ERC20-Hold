# ERC20-Hold
ERC20 token supporting balances on hold<br />
<br />
REQUIREMENTS:<br />
1-Install NodeJS v8.9.4 or later<br />
2-Install Truffle: 'npm install -g truffle'<br />
<br />
ERC20-HOLDABLE tests:<br />
1-On a cmd go to the project folder and run: "truffle develop"<br />
2-Run: "test" -> 11 tests on the ERC20-Holdable token should pass<br />
<br />
BETTING-GAME test:<br />
1-On a cmd go to the project folder and run: "truffle develop"<br />
2-Run: "migrate" and wait for the migration to finish.<br />
3-On a new cmd got o the project folder and run: "node index.js" and wait for "API escuchando en el puerto 8080" message<br />
4-Import the BETGAME.postman_collection in postman and run 4 times the 'Add player' request.<br />
5-Run all the 'bet' requests, for each one a message of 'bet created' should appear on the api cmd.<br />
6-On the 4th bet, we should see the winner on the api cmd.<br />
7-Run the 'Get user balance and bets' requests to check the users balances after the bet was executed.<br />
