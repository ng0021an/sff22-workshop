const RockSolidToken = artifacts.require("RockSolidToken");

module.exports = function (deployer) {
    deployer.deploy(RockSolidToken);
}