const ERC20Holdable = artifacts.require("ERC20Holdable");

module.exports = async function (deployer, network, accounts) {
  await deployer.deploy(ERC20Holdable, {from:accounts[0]});
  const deployedToken = await ERC20Holdable.deployed();
};
