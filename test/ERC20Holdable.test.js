const ERC20Holdable = artifacts.require('ERC20Holdable');
const truffleAssert = require('truffle-assertions');

contract("ERC20Holdable", (accounts) => {

  console.log(accounts);
  before( async () => {
    owner = accounts[1];
//    erc20holdable = await ERC20Holdable.at(owner);
    erc20holdable = await ERC20Holdable.deployed({from:owner});
})

  it("gives the owner 1000 tokens", async () => {
    let balance = await erc20holdable.balanceOf(owner);
    balance = web3.utils.fromWei(balance, 'ether');
    assert.equal(balance, '1000', 'balance should be 1M for the contract owner');
  })

  it('can transfer tokens between accounts', async () => {

    let balance = await erc20holdable.balanceOf(accounts[2])
    balance = web3.utils.fromWei(balance, 'ether');
    assert.equal(balance, '0', 'balance should be 0 for account 2');

    let amount = web3.utils.toWei('999', 'ether')
    await erc20holdable.transfer(accounts[2], amount , {from: owner })

    balance = await erc20holdable.balanceOf(accounts[2])
    balance = web3.utils.fromWei(balance, 'ether');
    assert.equal(balance, '999', 'balance should be 1000 for account 2');
  })

  it('event hold', async () => {

    let senderBalance = await erc20holdable.balanceOf(accounts[2])
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '999', 'balance should be 999 for sender before hold');

    let receiverBalance = await erc20holdable.balanceOf(accounts[3])
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '0', 'balance should be 0 for receiver before hold');

    let amount = web3.utils.toWei('500', 'ether')
    await erc20holdable.hold(1, accounts[3], amount , {from: accounts[2] })

    senderBalance = await erc20holdable.balanceOf(accounts[2])
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '499', 'balance should be 499 for sender after hold');

    receiverBalance = await erc20holdable.balanceOf(accounts[3])
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '0', 'balance should be 0 for receiver after hold');
  })

  it('event hold executed', async () => {

    let senderBalance = await erc20holdable.balanceOf(accounts[2]);
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '499', 'balance should be 499 for sender before hold executed');

    let receiverBalance = await erc20holdable.balanceOf(accounts[3]);
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '0', 'balance should be 0 for receiver before hold executed');

    await erc20holdable.executeHold(1);

    receiverBalance = await erc20holdable.balanceOf(accounts[3]);
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '500', 'balance should be 499 for receiver after hold executed');
  })

  it('event hold reverted', async () => {

    let senderBalance = await erc20holdable.balanceOf(accounts[2]);
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '499', 'balance should be 999 for sender before hold');

    let receiverBalance = await erc20holdable.balanceOf(accounts[3]);
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '500', 'balance should be 0 for receiver before hold');

    let amount = web3.utils.toWei('100', 'ether');
    await erc20holdable.hold(2, accounts[3], amount , {from: accounts[2] });

    senderBalance = await erc20holdable.balanceOf(accounts[2]);
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '399', 'balance should be 399 for sender after hold');

    receiverBalance = await erc20holdable.balanceOf(accounts[3]);
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '500', 'balance should be 500 for receiver after hold');

    await erc20holdable.removeHold(2);

    senderBalance = await erc20holdable.balanceOf(accounts[2]);
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '499', 'balance should be 499 for sender after hold removed');

    receiverBalance = await erc20holdable.balanceOf(accounts[3]);
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '500', 'balance should be 500 for receiver after hold removed');
  })

  it('HoldId must be unique', async() => {
    let amount = web3.utils.toWei('100', 'ether');
    await erc20holdable.hold(1, accounts[3], amount , {from: accounts[2] })
    await truffleAssert.reverts(erc20holdable.hold(1, accounts[3], amount , {from: accounts[2] }));
  })

  it('Hold does not exist for executeHold', async() => {
    await truffleAssert.reverts(erc20holdable.executeHold(4));
  })

  it('Hold does not exist for removeHold', async() => {
    await truffleAssert.reverts(erc20holdable.removeHold(4));
  })

})
