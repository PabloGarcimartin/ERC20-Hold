var express = require('express')
var app = express()
var erc20Holdable = require('./bettingGame.js')

var port = process.env.PORT || 8080

//Create player
app.post('/player', function(req, res) {

  erc20Holdable.addPlayer().then(result => {
      res.send(result);
  }).catch(err => {
      console.log(err);
      res.sendStatus(501);
  });
})

//Get all bets
app.get('/bets', function(req, res) {
  erc20Holdable.getBets().then(result => {
      res.send(result);
  }).catch(err => {
      console.log(err);
      res.sendStatus(501);
  });
})

//Get user balance and bets
app.get('/player/:address', function(req, res) {
  if (!req.params.address){
    res.status(501).send('User address is required');
    return;
  }

  erc20Holdable.playerInfo(req.params.address).then(result => {
      res.send(result);
  }).catch(err => {
      console.log(err);
      res.sendStatus(501);
  });
})

//Bet
app.post('/bet/:address', function(req, res) {
  if (!req.params.address){
    res.status(501).send('User address is required');
    return;
  }

  erc20Holdable.bet(req.params.address).then(result => {
      res.send(result);
  }).catch(err => {
      console.log(err);
      res.sendStatus(501);
  });
})

app.listen(port)
console.log('API escuchando en el puerto ' + port);
