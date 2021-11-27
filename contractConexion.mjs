import Web3 from 'web3';
import fs from 'fs';

let web3 = new Web3(new Web3.providers.HttpProvider('HTTP://127.0.0.1:9545'));
let contractAbi = JSON.parse(fs.readFileSync('./build/contracts/ERC20Holdable.json', 'utf8'));
let contractAddress = contractAbi.networks['5777'].address;

let contract = new web3.eth.Contract(contractAbi.abi, contractAddress);

export {contract};
