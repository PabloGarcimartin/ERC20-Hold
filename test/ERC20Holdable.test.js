const ERC20Holdable = artifacts.require('ERC20Holdable');
const truffleAssert = require('truffle-assertions');

contract("ERC20Holdable", (accounts) => {
  before( async () => {
    owner = accounts[0];
    erc20holdable = await ERC20Holdable.deployed({from:owner});
})

  it("gives the owner 10000 tokens", async () => {
    let balance = await erc20holdable.balanceOf(owner);
    balance = web3.utils.fromWei(balance, 'ether');
    assert.equal(balance, '10000', 'balance should be 1M for the contract owner');
  })

  it('can transfer tokens between accounts', async () => {

    let balance = await erc20holdable.balanceOf(accounts[1]);
    balance = web3.utils.fromWei(balance, 'ether');
    assert.equal(balance, '0', 'balance should be 0 for account 1');

    let amount = web3.utils.toWei('999', 'ether');
    await erc20holdable.transfer(accounts[1], amount , {from: owner });

    balance = await erc20holdable.balanceOf(accounts[1]);
    balance = web3.utils.fromWei(balance, 'ether');
    assert.equal(balance, '999', 'balance should be 999 for account 1');

    await erc20holdable.transfer(owner, amount , {from: accounts[1] });

    balance = await erc20holdable.balanceOf(accounts[1]);
    balance = web3.utils.fromWei(balance, 'ether');
    assert.equal(balance, '0', 'balance should be 0 for account 1');
  })

  it("can transferFrom tokens between accounts", async() => {

    let amount = web3.utils.toWei('1000', 'ether');
    await erc20holdable.transfer(accounts[1], amount , {from: owner });

    let allowanceAmount = web3.utils.toWei('500', 'ether');
    await erc20holdable.approve(accounts[1], amount);

    let allowance = await erc20holdable.allowance(accounts[0], accounts[1]);
    allowance = web3.utils.fromWei(allowance, 'ether');
    assert.equal(allowance, '1000', 'allowance is wrong');

    let smallAmount = web3.utils.toWei('200', 'ether');
    await erc20holdable.approve(accounts[1], smallAmount);
    await erc20holdable.transferFrom(accounts[0], accounts[2], smallAmount, {from: accounts[1]});

    let balance0 = await erc20holdable.balanceOf(accounts[0]);
    balance0 = web3.utils.fromWei(balance0, 'ether');
    assert.equal(balance0, '8800', 'accounts[0] balance is wrong');

    let balance1 = await erc20holdable.balanceOf(accounts[1]);
    balance1 = web3.utils.fromWei(balance1, 'ether');
    assert.equal(balance1, '1000', 'accounts[1] balance is wrong');

    let balance2 = await erc20holdable.balanceOf(accounts[2]);
    balance2 = web3.utils.fromWei(balance2, 'ether');
    assert.equal(balance2, '200', 'accounts[2] balance is wrong');

    await erc20holdable.transfer(owner, smallAmount , {from: accounts[2] });
    balance2 = await erc20holdable.balanceOf(accounts[2]);
    balance2 = web3.utils.fromWei(balance2, 'ether');
    assert.equal(balance2, '0', 'accounts[2] balance is wrong');
  });

  it('event hold', async () => {

    let amount = web3.utils.toWei('999', 'ether')
    await erc20holdable.transfer(accounts[2], amount , {from: owner })

    let senderBalance = await erc20holdable.balanceOf(accounts[2])
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '999', 'balance should be 999 for sender before hold');

    let receiverBalance = await erc20holdable.balanceOf(accounts[3])
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '0', 'balance should be 0 for receiver before hold');

    let holdAmount = web3.utils.toWei('500', 'ether')
    await erc20holdable.hold(1, accounts[3], holdAmount , {from: accounts[2] })

    senderBalance = await erc20holdable.balanceOf(accounts[2])
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '499', 'balance should be 499 for sender after hold');

    receiverBalance = await erc20holdable.balanceOf(accounts[3])
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '0', 'balance should be 0 for receiver after hold');
  })

  it('event hold executed', async () => {

    let amount = web3.utils.toWei('999', 'ether')
    await erc20holdable.transfer(accounts[4], amount , {from: owner })

    let senderBalanceBeforeHold = await erc20holdable.balanceOf(accounts[4]);
    senderBalanceBeforeHold = web3.utils.fromWei(senderBalanceBeforeHold, 'ether');
    assert.equal(senderBalanceBeforeHold, '999', 'balance should be 999 for sender before hold');

    let receiverBalanceBeforeHold = await erc20holdable.balanceOf(accounts[5]);
    receiverBalanceBeforeHold = web3.utils.fromWei(receiverBalanceBeforeHold, 'ether');
    assert.equal(receiverBalanceBeforeHold, '0', 'balance should be 0 for receiver before hold');

    let holdAmount = web3.utils.toWei('500', 'ether')
    await erc20holdable.hold(2, accounts[5], holdAmount , {from: accounts[4] })

    let senderBalanceBeforeHoldExecuted = await erc20holdable.balanceOf(accounts[4])
    senderBalanceBeforeHoldExecuted = web3.utils.fromWei(senderBalanceBeforeHoldExecuted, 'ether');
    assert.equal(senderBalanceBeforeHoldExecuted, '499', 'balance should be 499 for sender before hold executed');

    let receiverBalanceBeforeHoldExecuted = await erc20holdable.balanceOf(accounts[5])
    receiverBalanceBeforeHoldExecuted = web3.utils.fromWei(receiverBalanceBeforeHoldExecuted, 'ether');
    assert.equal(receiverBalanceBeforeHoldExecuted, '0', 'balance should be 0 for receiver before hold executed');

    await erc20holdable.executeHold(2, {from: accounts[4]});

    let receiverBalanceAfterHoldExecuted = await erc20holdable.balanceOf(accounts[5]);
    receiverBalanceAfterHoldExecuted = web3.utils.fromWei(receiverBalanceAfterHoldExecuted, 'ether');
    assert.equal(receiverBalanceAfterHoldExecuted, '500', 'balance should be 500 for receiver after hold executed');
  })

  it('event hold canceled', async () => {

    let amount = web3.utils.toWei('999', 'ether')
    await erc20holdable.transfer(accounts[6], amount , {from: owner })

    let senderBalance = await erc20holdable.balanceOf(accounts[6]);
    senderBalance = web3.utils.fromWei(senderBalance, 'ether');
    assert.equal(senderBalance, '999', 'balance should be 999 for sender before hold');

    let receiverBalance = await erc20holdable.balanceOf(accounts[7]);
    receiverBalance = web3.utils.fromWei(receiverBalance, 'ether');
    assert.equal(receiverBalance, '0', 'balance should be 0 for receiver before hold');

    let holdAmount = web3.utils.toWei('100', 'ether');
    await erc20holdable.hold(3, accounts[7], holdAmount , {from: accounts[6] });

    senderBalanceBeforeHoldCanceled = await erc20holdable.balanceOf(accounts[6]);
    senderBalanceBeforeHoldCanceled = web3.utils.fromWei(senderBalanceBeforeHoldCanceled, 'ether');
    assert.equal(senderBalanceBeforeHoldCanceled, '899', 'balance should be 399 for sender after hold');

    receiverBalanceBeforeHoldCanceled = await erc20holdable.balanceOf(accounts[7]);
    receiverBalanceBeforeHoldCanceled = web3.utils.fromWei(receiverBalanceBeforeHoldCanceled, 'ether');
    assert.equal(receiverBalanceBeforeHoldCanceled, '0', 'balance should be 0 for receiver after hold');

    await erc20holdable.cancelHold(3, {from: accounts[0]});

    senderBalanceAfterHoldCanceled = await erc20holdable.balanceOf(accounts[6]);
    senderBalanceAfterHoldCanceled = web3.utils.fromWei(senderBalanceAfterHoldCanceled, 'ether');
    assert.equal(senderBalanceAfterHoldCanceled, '999', 'balance should be 499 for sender after hold canceled');

    receiverBalanceAfterHoldCanceled = await erc20holdable.balanceOf(accounts[7]);
    receiverBalanceAfterHoldCanceled = web3.utils.fromWei(receiverBalanceAfterHoldCanceled, 'ether');
    assert.equal(receiverBalanceAfterHoldCanceled, '0', 'balance should be 0 for receiver after hold canceled');
  })

  it('event hold from', async () => {

    let amount = web3.utils.toWei('1000', 'ether');
    await erc20holdable.transfer(accounts[8], amount , {from: owner });

    let allowanceAmount = web3.utils.toWei('500', 'ether');
    await erc20holdable.approve(accounts[8], amount);

    let allowance = await erc20holdable.allowance(owner, accounts[8]);
    allowance = web3.utils.fromWei(allowance, 'ether');
    assert.equal(allowance, '1000', 'allowance is wrong');

    let smallAmount = web3.utils.toWei('200', 'ether');
    await erc20holdable.approve(accounts[8], smallAmount);
    await erc20holdable.holdFrom(3, owner, accounts[9], smallAmount, {from: accounts[8]});

    // let balance0 = await erc20holdable.balanceOf(accounts[7]);
    // balance0 = web3.utils.fromWei(balance0, 'ether');
    // assert.equal(balance0, '8800', 'accounts[7] balance is wrong');

    let balance1 = await erc20holdable.balanceOf(accounts[8]);
    balance1 = web3.utils.fromWei(balance1, 'ether');
    assert.equal(balance1, '1000', 'accounts[8] balance is wrong');

    let balance2 = await erc20holdable.balanceOf(accounts[9]);
    balance2 = web3.utils.fromWei(balance2, 'ether');
    assert.equal(balance2, '0', 'accounts[9] balance is wrong');

    await erc20holdable.executeHold(3, {from: owner});

    balance1 = await erc20holdable.balanceOf(accounts[8]);
    balance1 = web3.utils.fromWei(balance1, 'ether');
    assert.equal(balance1, '1000', 'accounts[8] balance is wrong');

    balance2 = await erc20holdable.balanceOf(accounts[9]);
    balance2 = web3.utils.fromWei(balance2, 'ether');
    assert.equal(balance2, '200', 'accounts[9] balance is wrong');

    await erc20holdable.transfer(owner, smallAmount , {from: accounts[9] });
    balance2 = await erc20holdable.balanceOf(accounts[9]);
    balance2 = web3.utils.fromWei(balance2, 'ether');
    assert.equal(balance2, '0', 'accounts[9] balance is wrong');
  })

  it('Hold Id must be unique', async() => {
    let amount = web3.utils.toWei('100', 'ether');
    await erc20holdable.hold(4, accounts[3], amount , {from: accounts[2] })
    await truffleAssert.reverts(erc20holdable.hold(4, accounts[3], amount , {from: accounts[2] }));
  })

  it('Hold not executed by contract owner or hold creator', async() => {
    let amount = web3.utils.toWei('100', 'ether');
    await erc20holdable.hold(5, accounts[3], amount , {from: accounts[2] })
    await truffleAssert.reverts(erc20holdable.executeHold(5, {from: accounts[4] }));
  })

  it('Hold not canceled by contract owner', async() => {
    let amount = web3.utils.toWei('100', 'ether');
    await erc20holdable.hold(6, accounts[3], amount , {from: accounts[2] });
    await truffleAssert.reverts(erc20holdable.cancelHold(6, {from: accounts[2] }));
    await truffleAssert.reverts(erc20holdable.cancelHold(6, {from: accounts[4] }));
  })

  it('Hold Id does not exist for executeHold()', async() => {
    await truffleAssert.reverts(erc20holdable.executeHold(7));
  })

  it('Hold Id does not exist for cancelHold()', async() => {
    await truffleAssert.reverts(erc20holdable.cancelHold(7));
  })

})
