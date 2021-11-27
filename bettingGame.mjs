import Web3 from 'web3';
import {contract} from './contractConexion.mjs';
const web3 = new Web3(new Web3.providers.HttpProvider('HTTP://127.0.0.1:9545'));

const accounts = await web3.eth.getAccounts();
const owner = accounts[0];

let players = [];
let openBets = [];
let betsByUser = [];
let totalBets = 0;
const _initialAmount = 100;
const _betImport = 5;
const _maxBets = 4;
const _gasLimit = web3.utils.toWei('1');

var bettingGame = {
            addPlayer: async function () {
              if( players.length == accounts.length - 1) {
                let newAccount = await web3.eth.accounts.create();
                accounts.push(newAccount.address);
              }
              let playerAccount = accounts[players.length + 1];

              await contract.methods.transfer(playerAccount, web3.utils.toWei(_initialAmount.toString())).send({from: owner}).then(result => {
                players.push(playerAccount);
                betsByUser[playerAccount] = [];
              }).catch(err => {
                console.log('Error: ' + err);
                return err;
              });

              return await this.playerInfo(playerAccount);
            },
            playerInfo: async function (address) {
              if( !players.includes(address) ){
                return { error: 'Address doesn´t match any player'};
              }

              let userBalance =  web3.utils.fromWei(await contract.methods.balanceOf(address).call({from: owner}));
              let userInfo = {
                address: address,
                balance: userBalance,
                bets: betsByUser[address]
              };
              return userInfo;
            },
            bet: async function (address) {
              if( !players.includes(address) ){
                return { error: 'Address doesn´t match any player'};
              }

              let userBalance = web3.utils.fromWei(await contract.methods.balanceOf(address).call({from: address}));
              if( userBalance < _betImport ){
                return { error: 'Not enough tokens to bet'};
              }

              let bet;
              try {
                let estimatedGas = await contract.methods.hold(totalBets, owner, web3.utils.toWei( _betImport.toString() )).estimateGas({from: address});

                await contract.methods.hold(totalBets,owner, web3.utils.toWei( _betImport.toString() )).send({from: address, gas: estimatedGas})
                .then( result => {
                  if( result.status ){
                    console.log('Bet created');
                    betsByUser[address].push([totalBets,'pending']);
                    openBets.push([totalBets, address]);
                    totalBets = totalBets + 1;
                  }
                }).catch(err => {
                  console.log('Error: ' + err);
                  return err;
                });
              } catch (error) {
                console.log('Error: ' + error);
                return error;
              }

              if( openBets.length == _maxBets ){
                console.log('WE HAVE 4 BETS!');
                this.runBets();
              }

            },
            runBets: async function () {
              console.log(openBets);

              for( let i = 0; i<_maxBets; i++ ){
                  let betId = openBets[i][0];
                  let betAddress = openBets[i][1];
                  let estimatedGas = await contract.methods.executeHold(betId).estimateGas({from:owner});
                  await contract.methods.executeHold(betId).send({from: owner, gas: estimatedGas}).then(result => {
                    console.log('HOLD EXECUTED');
                    for( let j = 0; j < betsByUser[openBets[i][1]].length; j++  ){
                      if( betsByUser[betAddress][j][0] == betId ){
                        betsByUser[betAddress][j][1] = 'lost';
                      }
                    }
                  }).catch(err => {
                    console.log('Error: ' + err);
                    return err;
                  });
              }

              let winningBetIndex = Math.floor(Math.random() * _maxBets) + 1;
              let winningBetId = openBets[winningBetIndex][0];
              let winner =  openBets[winningBetIndex][1];

              console.log('WINNER IS: ');
              console.log(winner);

              let res = await contract.methods.transfer(winner, web3.utils.toWei((_maxBets*_betImport).toString())).send({from: owner}).then(result => {
                console.log('WINNER MONEY TRANSFERRED');
                for( let i = 0; i < betsByUser[winner].length; i++  ){
                  if(betsByUser[winner][i][0] == winningBetId){
                    betsByUser[winner][i][1] = 'won';
                  }
                }
              }).catch(err => {
                console.log('Error: ' + err);
                return err;
              });
              openBets = [];
              return true;
            },
            changeBetStatus: async function (betId, status) {

            }
}

export {bettingGame}
