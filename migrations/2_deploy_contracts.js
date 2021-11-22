const ERC20Holdable = artifacts.require("ERC20Holdable");

module.exports = async function (deployer, network, accounts) {
  const _name = "ERC20Holdable";
  const _symbol = "ERC20H";
  const _decimals = 1;
  const _totalSupply = 5000;

  await deployer.deploy(ERC20Holdable, {from:accounts[1]});
  const deployedToken = await ERC20Holdable.deployed();
};
