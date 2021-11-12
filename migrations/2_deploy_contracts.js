const ERC20Holdable = artifacts.require("ERC20Holdable");

module.exports = function (deployer, network, accounts) {
  deployer.deploy(ERC20Holdable, 1000, {from:accounts[1]});
};
