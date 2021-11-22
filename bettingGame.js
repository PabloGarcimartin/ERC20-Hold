const path = require('path');
const fs = require('fs');
var Web3 = require('web3');

let web3 = new Web3(new Web3.providers.HttpProvider('HTTP://127.0.0.1:9545'))
let contractAbi = JSON.parse(fs.readFileSync('./build/contracts/ERC20Holdable.json', 'utf8'))
let contractAddress = contractAbi.networks['5777'].address;

// Get address
let players = [];
let accountsArray = [];
let bets = [];
const _initialAmount = '100';
const _betImport = '5';
const _maxBets = 4;
const _gasLimit = web3.utils.toWei('1');


var run = {
            addPlayer: async function () {
              let contract = await new web3.eth.Contract(contractAbi.abi, contractAddress);
              if(accountsArray.length == 0){
                accountsArray = await web3.eth.getAccounts();
              }
              let owner = accountsArray[1];

              if( players.length == accountsArray.length - 2) {
                let newAccount = await web3.eth.accounts.create();
                accountsArray.push(newAccount.address)
              }
              let playerAccount = accountsArray[players.length + 2];

              let res = await contract.methods.transfer(playerAccount, web3.utils.toWei(_initialAmount)).send({from: owner}).then(result => {
                players.push(playerAccount);
              }).catch(err => {
                console.log('Error: ' + err);
                return err;
              });

              let userData = await this.playerInfo(playerAccount);
              return userData;
            },
            playerInfo: async function (address) {
              if( !players.includes(address) ){
                return { error: 'Address doesn´t match any player'}
              }
              let contract = await new web3.eth.Contract(contractAbi.abi, contractAddress);
              let accounts = await web3.eth.getAccounts();
              let owner = accounts[1]

              let userTokens = await contract.methods.balanceOf(address).call({from: owner})
              let userBets =  await contract.methods.getUserHolds(address).call({from: owner})
              let userInfo = {
                address: address,
                balance: web3.utils.fromWei(userTokens),
                bets: userBets
              }
              return userInfo;
            },
            getBets: async function () {
              let contract = await new web3.eth.Contract(contractAbi.abi, contractAddress);
              let accounts = await web3.eth.getAccounts();
              let owner = accounts[1]

              let allBets =  await contract.methods.getHolds().call({from: owner})
              return allBets;
            },
            bet: async function (address) {
              if( !players.includes(address) ){
                return { error: 'Address doesn´t match any player'}
              }

              let contract = await new web3.eth.Contract(contractAbi.abi, contractAddress);
              let accounts = await web3.eth.getAccounts();
              let owner = accounts[1]

              let userTokens = web3.utils.fromWei(await contract.methods.balanceOf(address).call({from: address}))
              if( userTokens < parseInt(_betImport) ){
                return { error: 'Not enough tokens to bet'}
              }

              let bet;
              try {
                let oldBets = await contract.methods.getHolds().call({from: owner});
                let gas = await contract.methods.hold(oldBets.length,owner,web3.utils.toWei(_betImport)) .estimateGas({from: address})
                let res = await contract.methods.hold(oldBets.length,owner, web3.utils.toWei(_betImport)).send({from: address, gas: gas})
                .then(result => {
                  if(result.status){
                    console.log('Bet created');
                    bets.push(oldBets.length);
                  }
                }).catch(err => {
                  console.log('Error: ' + err);
                  return err;
                });
              } catch (e) {
                console.log('Error: ' + e);
                return e;
              }

              if( bets.length == _maxBets ){
                console.log('WE HAVE 4 BETS!');
                this.runBets(contract, owner);
              }

            },
            runBets: async function (contract, owner) {
              console.log(bets);
              let winningBet = Math.floor(Math.random() * _maxBets) + 1;
              let accounts = await web3.eth.getAccounts();

              for(let i = 0; i<_maxBets; i++){
                  if( i == winningBet ){
                    let winningBet = await contract.methods.getHold(bets[i]).call({from: owner});
                    let winner = winningBet['sender'];
                    console.log('WINNER IS: ');
                    console.log(winner);
                    let gas = await contract.methods.cancelHold(bets[i]).estimateGas({from:owner});
                    await contract.methods.cancelHold(bets[i]).send({from: owner, gas: gas}).then(result => {
                      console.log('WINNER HOLD REVERTED');
                    }).catch(err => {
                      console.log('Error: ' + err);
                      return err;
                    });;

                    let res = await contract.methods.transfer(winner, web3.utils.toWei('15')).send({from: owner}).then(result => {
                      console.log('WINNER MONEY TRANSFERRED');
                    }).catch(err => {
                      console.log('Error: ' + err);
                      return err;
                    });
                  } else {
                    let gas = await contract.methods.executeHold(bets[i]).estimateGas({from:owner});
                    await contract.methods.executeHold(bets[i]).send({from: owner, gas: gas}).then(result => {
                      console.log('LOSER HOLD EXECUTED');
                    }).catch(err => {
                      console.log('Error: ' + err);
                      return err;
                    });
                  }
              }
              bets = [];
              return true;
            }
}

module.exports = run
