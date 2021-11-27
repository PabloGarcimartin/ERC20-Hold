import {bettingGame} from './bettingGame.mjs';
import express from 'express';
var app = express();

var port = process.env.PORT || 8080;

//Create player
app.post('/players', function(req, res) {

  bettingGame.addPlayer().then(result => {
      res.status(201).send(result);
  }).catch(err => {
      console.log(err);
      res.sendStatus(500);
  });
})

//Bet
app.post('/players/:address/bets', function(req, res) {
  if (!req.params.address){
    res.status(400).send('User address is required');
    return;
  }

  bettingGame.bet(req.params.address).then(result => {
      res.status(201).send(result);
  }).catch(err => {
      console.log(err);
      res.sendStatus(500);
  });
})

//Get user balance and bets
app.get('/players/:address', function(req, res) {
  if (!req.params.address){
    res.status(400).send('User address is required');
    return;
  }

  bettingGame.playerInfo(req.params.address).then(result => {
      res.send(result);
  }).catch(err => {
      console.log(err);
      res.sendStatus(500);
  });
})


app.listen(port);
console.log('API escuchando en el puerto ' + port);
